package sys

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
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
	menuQueryErr    error
	statusQueryErr  error
	unchangedRows   bool
}

func (d *genCodesStatusDriver) DoSelect(ctx context.Context, link gdb.Link, query string, args ...any) (gdb.Result, error) {
	fixture := ctx.Value(genCodesUpdateKey{}).(*genCodesStatusFixture)
	if strings.Contains(query, "admin_menu") {
		return nil, fixture.menuQueryErr
	}
	if fixture.statusQueryErr != nil {
		return nil, fixture.statusQueryErr
	}
	if fixture.record == nil {
		return nil, nil
	}
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
	if fixture.record == nil {
		return driver.RowsAffected(0), nil
	}
	for key, value := range update {
		fixture.record[key] = value
	}
	if fixture.unchangedRows {
		return driver.RowsAffected(0), nil
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
	disableGenCodesFileLogs(t)
	newContext := func() (context.Context, *genCodesStatusFixture) {
		fixture := &genCodesStatusFixture{record: map[string]any{"id": int64(52), "status": consts.GenCodesStatusOk, "gen_type": consts.GenCodesTypeQueue, "var_name": "AppLog"}}
		return context.WithValue(context.Background(), genCodesUpdateKey{}, fixture), fixture
	}
	newInput := func() entity.SysGenCodes {
		return entity.SysGenCodes{Id: 52, GenType: consts.GenCodesTypeQueue, VarName: "AppLog", Status: consts.GenCodesStatusWait}
	}

	t.Run("stage logs preserve errors and panic semantics", func(t *testing.T) {
		logger := g.Log()
		original := logger.GetConfig()
		defer func() {
			if err := logger.SetConfig(original); err != nil {
				t.Error(err)
			}
		}()
		var output bytes.Buffer
		config := original
		config.Path, config.Writer, config.StdoutPrint = "", &output, false
		config.Level, config.Flags = glog.LEVEL_ALL, config.Flags&^glog.F_ASYNC
		if err := logger.SetConfig(config); err != nil {
			t.Fatal(err)
		}
		for _, scenario := range []string{"success", "failure", "panic"} {
			output.Reset()
			ctx, _ := newContext()
			wantErr := errors.New("stage failure")
			var stageErr error
			panicErr := g.Try(ctx, func(ctx context.Context) {
				stageErr = runBuildStage(ctx, 52, "test-stage", func(stageCtx context.Context) error {
					if stageCtx != ctx {
						t.Error("stage context was replaced")
					}
					switch scenario {
					case "failure":
						return wantErr
					case "panic":
						panic("stage panic")
					}
					return nil
				})
			})
			if scenario == "failure" && stageErr != wantErr || scenario == "panic" && panicErr == nil || scenario == "success" && (stageErr != nil || panicErr != nil) {
				t.Fatalf("stage changed outcome for %s: error=%v panic=%v", scenario, stageErr, panicErr)
			}
			text := output.String()
			for _, value := range []string{"代码生成阶段开始", "代码生成阶段结束", "52", "test-stage", "elapsedMs", "succeeded"} {
				if !strings.Contains(text, value) {
					t.Errorf("stage %s log is missing %q: %s", scenario, value, text)
				}
			}
		}
	})

	t.Run("missing status record is rejected", func(t *testing.T) {
		ctx, fixture := newContext()
		fixture.record = nil
		err := NewSysGenCodes().Status(ctx, &sysin.GenCodesStatusInp{Id: 52, Status: consts.GenCodesStatusOk})
		if err == nil || !strings.Contains(err.Error(), "生成配置不存在") {
			t.Fatalf("missing record reported as saved: %v", err)
		}
	})
	t.Run("unchanged status with zero affected rows succeeds", func(t *testing.T) {
		ctx, fixture := newContext()
		fixture.unchangedRows = true
		if err := NewSysGenCodes().Status(ctx, &sysin.GenCodesStatusInp{Id: 52, Status: consts.GenCodesStatusOk}); err != nil {
			t.Fatalf("same status treated as missing record: %v", err)
		}
	})
	t.Run("status existence query errors are returned", func(t *testing.T) {
		ctx, fixture := newContext()
		fixture.unchangedRows = true
		fixture.statusQueryErr = errors.New("status lookup unavailable")
		if err := NewSysGenCodes().Status(ctx, &sysin.GenCodesStatusInp{Id: 52, Status: consts.GenCodesStatusOk}); err == nil {
			t.Fatal("status lookup failure hidden")
		}
	})
	t.Run("missing menu cleanup is retryable", func(t *testing.T) {
		ctx, _ := newContext()
		for attempt := 0; attempt < 2; attempt++ {
			if err := NewSysGenCodes().deleteBuildMenus(ctx, "appLog"); err != nil {
				t.Fatalf("absent menu prevents cleanup retry: %v", err)
			}
		}
	})
	t.Run("menu query errors are returned", func(t *testing.T) {
		ctx, fixture := newContext()
		fixture.menuQueryErr = errors.New("menu lookup unavailable")
		if err := NewSysGenCodes().deleteBuildMenus(ctx, "appLog"); err == nil {
			t.Fatal("real menu lookup failure hidden")
		}
	})
	t.Run("cleanup can retry after service refresh fails", func(t *testing.T) {
		t.Chdir(t.TempDir())
		adapter, err := gcfg.NewAdapterContent(`{"hggen":{"application":{"crud":{"templates":[{
			"apiPath":"generated/api", "inputPath":"generated/input", "controllerPath":"generated/controller",
			"logicPath":"internal/logic/sys", "routerPath":"generated/router", "sqlPath":"generated/sql",
			"webApiPath":"generated/web/api", "webViewsPath":"generated/web/views"
		}]}}}}`)
		if err != nil {
			t.Fatal(err)
		}
		oldAdapter := g.Cfg().GetAdapter()
		g.Cfg().SetAdapter(adapter)
		defer g.Cfg().SetAdapter(oldAdapter)
		ctx, fixture := newContext()
		fixture.record["gen_type"] = consts.GenCodesTypeCurd
		fixture.record["status"] = consts.GenCodesStatusFail
		fixture.record["options"] = gjson.New(`{"autoOps":["genMenuPermissions"]}`)
		service := NewSysGenCodes()
		paths, _, err := service.getBuildFiles(ctx, &entity.SysGenCodes{VarName: "AppLog", Options: fixture.record["options"].(*gjson.Json)})
		if err != nil {
			t.Fatal(err)
		}
		for _, path := range paths {
			if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				t.Fatal(err)
			}
			if err = os.WriteFile(path, []byte("partial generated file"), 0644); err != nil {
				t.Fatal(err)
			}
		}
		// go.mod 暂缺，强制第一次在删除文件、清理缺失菜单之后刷新 service 失败。
		if err = service.DeleteBuild(ctx, &sysin.GenCodesDeleteBuildInp{Id: 52}); err == nil || !strings.Contains(err.Error(), "更新业务服务失败") {
			t.Fatalf("expected service refresh failure after cleanup: %v", err)
		}
		for _, path := range paths {
			if _, statErr := os.Stat(path); !os.IsNotExist(statErr) {
				t.Fatalf("partial file was not removed: %s %v", path, statErr)
			}
		}
		if err = os.WriteFile("go.mod", []byte("module example.com/cleanupretry\n\ngo 1.24.4\n"), 0644); err != nil {
			t.Fatal(err)
		}
		if err = service.DeleteBuild(ctx, &sysin.GenCodesDeleteBuildInp{Id: 52}); err != nil {
			t.Fatalf("cleanup retry failed: %v", err)
		}
		if got := gconv.Int(fixture.record["status"]); got != consts.GenCodesStatusWait {
			t.Fatalf("cleanup retry did not reset status: %d", got)
		}
		// 清理的状态重置也必须走显式事务，提交失败不能被报告为清理成功。
		fixture.record["status"] = consts.GenCodesStatusFail
		previousPool := genCodesSessionPool
		genCodesSessionPool = sql.OpenDB(genCodesSessionConnector{&genCodesSessionStore{failCommit: true}})
		defer func() {
			genCodesSessionPool.Close()
			genCodesSessionPool = previousPool
		}()
		if err = service.DeleteBuild(ctx, &sysin.GenCodesDeleteBuildInp{Id: 52}); err == nil || !strings.Contains(err.Error(), "更新生成状态失败") {
			t.Fatalf("cleanup status commit failure hidden: %v", err)
		}
	})

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
	t.Run("record removed during generation cannot report success", func(t *testing.T) {
		ctx, fixture := newContext()
		err := NewSysGenCodes().build(ctx, &sysin.GenCodesBuildInp{SysGenCodes: newInput()}, func(context.Context, *sysin.GenCodesBuildInp) error {
			fixture.record = nil
			return nil
		})
		if err == nil || !strings.Contains(err.Error(), "结果状态保存失败") {
			t.Fatalf("removed record reported as completed: %v", err)
		}
	})
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
