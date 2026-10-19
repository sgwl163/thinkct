package sys

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
)

type genCodesStatusDriver struct{ *gdb.DriverDefault }

func (d *genCodesStatusDriver) Master(...string) (*sql.DB, error) {
	return genCodesSessionPool, nil
}

func (d *genCodesStatusDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &genCodesStatusDriver{DriverDefault: &gdb.DriverDefault{Core: core}}, nil
}

func (d *genCodesStatusDriver) TableFields(ctx context.Context, table string, schema ...string) (map[string]*gdb.TableField, error) {
	fields := make(map[string]*gdb.TableField)
	for _, name := range []string{"id", "gen_type", "gen_template", "var_name", "options", "master_columns", "status", "created_at", "updated_at"} {
		fields[name] = &gdb.TableField{Name: name, Type: "text"}
	}
	return fields, nil
}

type genCodesUpdateKey struct{}
type genCodesStatusFixture struct {
	record          map[string]any
	updates         []map[string]any
	failFinalStatus bool
}

func (d *genCodesStatusDriver) DoSelect(ctx context.Context, link gdb.Link, query string, args ...any) (gdb.Result, error) {
	fixture := ctx.Value(genCodesUpdateKey{}).(*genCodesStatusFixture)
	record := make(gdb.Record)
	for key, value := range fixture.record {
		record[key] = gvar.New(value)
	}
	return gdb.Result{record}, nil
}

func (d *genCodesStatusDriver) DoInsert(ctx context.Context, link gdb.Link, table string, data gdb.List, option gdb.DoInsertOption) (sql.Result, error) {
	fixture := ctx.Value(genCodesUpdateKey{}).(*genCodesStatusFixture)
	fixture.record = gconv.Map(data[0])
	fixture.record["id"] = int64(53)
	return genCodesStatusResult{}, nil
}

func (d *genCodesStatusDriver) DoUpdate(ctx context.Context, link gdb.Link, table string, data any, condition string, args ...any) (sql.Result, error) {
	fixture := ctx.Value(genCodesUpdateKey{}).(*genCodesStatusFixture)
	update := gconv.Map(data)
	fixture.updates = append(fixture.updates, update)
	if fixture.failFinalStatus && gconv.Int(update["status"]) == consts.GenCodesStatusFail {
		return nil, errors.New("status storage unavailable")
	}
	for key, value := range update {
		fixture.record[key] = value
	}
	return genCodesStatusResult{}, nil
}

type genCodesStatusResult struct{}

func (genCodesStatusResult) LastInsertId() (int64, error) { return 53, nil }
func (genCodesStatusResult) RowsAffected() (int64, error) { return 1, nil }

// 独立进程隔离 GoFrame 的默认数据库单例，测试不会连接真实数据库或影响其他测试。
func TestGenCodesEditPreservesBuildStatus(t *testing.T) {
	const childEnv = "HOTGO_GEN_STATUS_TEST"
	if os.Getenv(childEnv) != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestGenCodesEditPreservesBuildStatus$", "-test.v")
		cmd.Env = append(os.Environ(), childEnv+"=1")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("status regressions failed: %v\n%s", err, output)
		}
		return
	}
	t.Chdir(t.TempDir())
	const driverName = "gencodes_status_test"
	if err := gdb.Register(driverName, &genCodesStatusDriver{}); err != nil {
		t.Fatal(err)
	}
	if err := gdb.SetConfigGroup("default", gdb.ConfigGroup{{Type: driverName}}); err != nil {
		t.Fatal(err)
	}
	g.DB("default")
	genCodesSessionPool = sql.OpenDB(genCodesSessionConnector{&genCodesSessionStore{}})
	defer genCodesSessionPool.Close()
	logConfig := g.Log().GetConfig()
	logConfig.Path = ""
	if err := g.Log().SetConfig(logConfig); err != nil {
		t.Fatal(err)
	}
	newContext := func() (context.Context, *genCodesStatusFixture) {
		fixture := &genCodesStatusFixture{record: map[string]any{"id": int64(52), "status": consts.GenCodesStatusOk, "gen_type": consts.GenCodesTypeQueue, "var_name": "AppLog"}}
		return context.WithValue(context.Background(), genCodesUpdateKey{}, fixture), fixture
	}
	newInput := func() entity.SysGenCodes {
		return entity.SysGenCodes{Id: 52, GenType: consts.GenCodesTypeQueue, VarName: "AppLog", Status: consts.GenCodesStatusWait}
	}

	t.Run("stale save preserves result", func(t *testing.T) {
		ctx, fixture := newContext()
		result, err := NewSysGenCodes().Edit(ctx, &sysin.GenCodesEditInp{SysGenCodes: newInput()})
		if err != nil {
			t.Fatal(err)
		}
		if len(fixture.updates) != 1 {
			t.Fatalf("expected one update, got %d", len(fixture.updates))
		}
		for _, key := range []string{"id", "status", "created_at"} {
			if value, exists := fixture.updates[0][key]; exists {
				t.Errorf("configuration save overwrites protected %s with %v", key, value)
			}
		}
		if result.Status != consts.GenCodesStatusOk {
			t.Fatalf("response contains stale status %d", result.Status)
		}
	})
	for _, scenario := range []struct {
		name                           string
		fail, panic, cancel, newRecord bool
	}{
		{name: "successful build"}, {name: "partial failure", fail: true}, {name: "panic", panic: true},
		{name: "client cancellation", cancel: true}, {name: "new configuration", newRecord: true},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			ctx, fixture := newContext()
			if scenario.cancel {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				cancel()
			}
			in := &sysin.GenCodesBuildInp{SysGenCodes: newInput()}
			if scenario.newRecord {
				in.Id = 0
			}
			err := NewSysGenCodes().build(ctx, in, func(ctx context.Context, in *sysin.GenCodesBuildInp) error {
				if ctx.Err() != nil {
					t.Error("client cancellation cancels generation")
				}
				if in.Id <= 0 {
					t.Error("build did not receive saved record ID")
				}
				if gconv.Int(fixture.record["status"]) != consts.GenCodesStatusRunning {
					t.Error("build is not marked running before file writes")
				}
				if scenario.panic {
					panic("generation panic")
				}
				if scenario.fail {
					return errors.New("service publish failed")
				}
				return nil
			})
			want := consts.GenCodesStatusOk
			if scenario.fail || scenario.panic {
				want = consts.GenCodesStatusFail
				if err == nil {
					t.Error("failure was swallowed")
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if got := gconv.Int(fixture.record["status"]); got != want {
				t.Fatalf("want status %d, got %d", want, got)
			}
			if scenario.newRecord && in.Id != 53 {
				t.Fatalf("new record ID not propagated: %d", in.Id)
			}
		})
	}
	t.Run("failure status persistence error", func(t *testing.T) {
		ctx, fixture := newContext()
		fixture.failFinalStatus = true
		err := NewSysGenCodes().build(ctx, &sysin.GenCodesBuildInp{SysGenCodes: newInput()}, func(context.Context, *sysin.GenCodesBuildInp) error { return errors.New("generation failed") })
		if err == nil || !strings.Contains(err.Error(), "状态保存失败") {
			t.Fatalf("status persistence error hidden: %v", err)
		}
	})
	t.Run("query never persists an inferred failure", func(t *testing.T) {
		ctx, fixture := newContext()
		fixture.record["status"] = consts.GenCodesStatusRunning
		row := &sysin.GenCodesViewModel{SysGenCodes: newInput()}
		row.Status = consts.GenCodesStatusRunning
		if err := NewSysGenCodes().inspectBuild(ctx, row); err != nil {
			t.Fatal(err)
		}
		if len(fixture.updates) != 0 || gconv.Int(fixture.record["status"]) != consts.GenCodesStatusRunning {
			t.Fatalf("view/list query persisted an inferred failure: %v", fixture.updates)
		}
		if row.Status != consts.GenCodesStatusRunning || row.Building || row.BuildSucceeded || row.BuildStatusLabel != "结果待确认" || row.BuildMessage == "" {
			t.Fatalf("unconfirmed result was presented as success/failure: %+v", row)
		}
	})
	t.Run("finished build is not overwritten by a stale row", func(t *testing.T) {
		for _, staleStatus := range []int{consts.GenCodesStatusRunning, consts.GenCodesStatusWait, consts.GenCodesStatusFail} {
			ctx, fixture := newContext()
			row := &sysin.GenCodesViewModel{SysGenCodes: newInput()}
			row.Status = staleStatus
			if err := NewSysGenCodes().inspectBuild(ctx, row); err != nil {
				t.Fatal(err)
			}
			if row.Status != consts.GenCodesStatusOk || !row.BuildSucceeded || len(fixture.updates) != 0 {
				t.Fatalf("completed result overwritten by stale status %d: %+v", staleStatus, row)
			}
		}
	})
	t.Run("concurrent build rejected before save", func(t *testing.T) {
		ctx, fixture := newContext()
		lock := generationLock(0)
		if ok, err := lock.TryLock(); err != nil || !ok {
			t.Fatalf("lock: %v %v", ok, err)
		}
		defer lock.Close()
		err := NewSysGenCodes().build(ctx, &sysin.GenCodesBuildInp{SysGenCodes: newInput()}, func(context.Context, *sysin.GenCodesBuildInp) error {
			t.Error("concurrent build executed")
			return nil
		})
		if err == nil || len(fixture.updates) != 0 {
			t.Fatalf("concurrent build changed record: err=%v updates=%v", err, fixture.updates)
		}
	})
	t.Run("file existence is independent of execution status", func(t *testing.T) {
		ctx, fixture := newContext()
		t.Chdir(t.TempDir())
		adapter, err := gcfg.NewAdapterContent(`{"hggen":{"application":{"crud":{"templates":[{
			"apiPath":"generated/api", "inputPath":"generated/input", "controllerPath":"generated/controller",
			"logicPath":"generated/logic", "routerPath":"generated/router", "sqlPath":"generated/sql",
			"webApiPath":"generated/web/api", "webViewsPath":"generated/web/views"
		}]}}}}`)
		if err != nil {
			t.Fatal(err)
		}
		oldAdapter := g.Cfg().GetAdapter()
		g.Cfg().SetAdapter(adapter)
		defer g.Cfg().SetAdapter(oldAdapter)
		// 未创建的目标目录不能退化为工作目录，误报同名的无关文件。
		if err = os.MkdirAll("applog", 0755); err != nil {
			t.Fatal(err)
		}
		if err = os.WriteFile("applog/applog.go", []byte("unrelated"), 0644); err != nil {
			t.Fatal(err)
		}
		row := &sysin.GenCodesViewModel{SysGenCodes: newInput()}
		row.GenType = consts.GenCodesTypeCurd
		service := NewSysGenCodes()
		if err = service.inspectBuild(ctx, row); err != nil {
			t.Fatal(err)
		}
		if row.BuildFiles != 0 {
			t.Fatalf("nonexistent target counted unrelated files: %d", row.BuildFiles)
		}
		paths, _, err := service.getBuildFiles(ctx, &row.SysGenCodes)
		if err != nil {
			t.Fatal(err)
		}
		for _, path := range paths[:3] {
			if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				t.Fatal(err)
			}
			if err = os.WriteFile(path, []byte("generated"), 0644); err != nil {
				t.Fatal(err)
			}
		}
		for _, status := range []int{consts.GenCodesStatusWait, consts.GenCodesStatusFail, consts.GenCodesStatusOk} {
			row = &sysin.GenCodesViewModel{SysGenCodes: newInput()}
			row.GenType, row.Status = consts.GenCodesTypeCurd, status
			fixture.record["status"] = status
			if err = service.inspectBuild(ctx, row); err != nil {
				t.Fatal(err)
			}
			if row.Status != status || row.BuildFiles != 3 || row.BuildSucceeded != (status == consts.GenCodesStatusOk) || row.BuildMessage == "" {
				t.Fatalf("files incorrectly imply a completed build: %+v", row)
			}
			if status != consts.GenCodesStatusOk && row.BuildStatusLabel != consts.GenCodesExistingFilesStatusNameMap[status] {
				t.Fatalf("existing files are shown as never generated: %+v", row)
			}
		}
	})
	t.Run("active build stays running", func(t *testing.T) {
		ctx, fixture := newContext()
		fixture.record["status"] = consts.GenCodesStatusRunning
		lock := generationLock(52)
		if ok, err := lock.TryLock(); err != nil || !ok {
			t.Fatalf("lock: %v %v", ok, err)
		}
		defer lock.Close()
		row := &sysin.GenCodesViewModel{SysGenCodes: newInput()}
		row.Status = consts.GenCodesStatusRunning
		if err := NewSysGenCodes().inspectBuild(ctx, row); err != nil {
			t.Fatal(err)
		}
		if !row.Building || row.Status != consts.GenCodesStatusRunning {
			t.Fatalf("active generation mistaken for interruption: %+v", row)
		}
	})
}
