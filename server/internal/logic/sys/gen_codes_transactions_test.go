package sys

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/consts"
	"hotgo/internal/model/input/sysin"
)

// 使用真正的 database/sql 连接与事务；仅数据库存储为内存，绝不连接业务数据库。
type genCodesSessionStore struct {
	status     int64
	menus      int
	failCommit bool
	statements []string
}
type genCodesSessionConnector struct{ store *genCodesSessionStore }

func (c genCodesSessionConnector) Connect(context.Context) (driver.Conn, error) {
	return &genCodesSessionConn{store: c.store}, nil
}
func (c genCodesSessionConnector) Driver() driver.Driver { return genCodesSessionSQLDriver{} }

type genCodesSessionSQLDriver struct{}

func (genCodesSessionSQLDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("use connector")
}

type genCodesSessionConn struct {
	store        *genCodesSessionStore
	pending      *int64
	pendingMenus int
	transaction  bool
}

func (c *genCodesSessionConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("unexpected prepare")
}
func (c *genCodesSessionConn) Close() error              { c.pending = nil; return nil }
func (c *genCodesSessionConn) Begin() (driver.Tx, error) { c.transaction = true; return c, nil }
func (c *genCodesSessionConn) Commit() error {
	c.transaction = false
	if c.store.failCommit {
		return errors.New("commit unavailable")
	}
	if c.pending != nil {
		c.store.status = *c.pending
		c.pending = nil
	}
	c.store.menus += c.pendingMenus
	c.pendingMenus = 0
	return nil
}
func (c *genCodesSessionConn) Rollback() error {
	c.pending = nil
	c.pendingMenus = 0
	c.transaction = false
	return nil
}
func (c *genCodesSessionConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	c.store.statements = append(c.store.statements, query)
	upper := strings.ToUpper(query)
	switch {
	case strings.HasPrefix(upper, "INSERT "):
		c.pendingMenus++
		return driver.RowsAffected(1), nil
	case strings.HasPrefix(upper, "SET @"):
		return driver.RowsAffected(0), nil
	case upper == "FAIL;":
		return nil, errors.New("menu rejected")
	}
	if !strings.HasPrefix(upper, "UPDATE ") {
		return nil, errors.New("unexpected SQL: " + query)
	}
	// 模拟菜单导入后的 AUTOCOMMIT=0：无显式提交时 UPDATE 只对本连接可见。
	for _, arg := range args {
		value := gconv.Int64(arg.Value)
		if value >= 1 && value <= 4 {
			c.pending = &value
			if value == c.store.status {
				return driver.RowsAffected(0), nil
			}
			return driver.RowsAffected(1), nil
		}
	}
	return nil, errors.New("missing status")
}
func (c *genCodesSessionConn) QueryContext(_ context.Context, query string, _ []driver.NamedValue) (driver.Rows, error) {
	status := c.store.status
	if c.pending != nil {
		status = *c.pending
	}
	fields := strings.SplitN(strings.ToLower(query), "from", 2)[0]
	return &genCodesSessionRows{status: status, idOnly: !strings.Contains(fields, "status")}, nil
}

type genCodesSessionRows struct {
	status int64
	idOnly bool
	read   bool
}

func (r *genCodesSessionRows) Columns() []string {
	if r.idOnly {
		return []string{"id"}
	}
	return []string{"status"}
}
func (*genCodesSessionRows) Close() error { return nil }
func (r *genCodesSessionRows) Next(values []driver.Value) error {
	if r.read {
		return io.EOF
	}
	r.read = true
	if r.idOnly {
		values[0] = int64(56)
	} else {
		values[0] = r.status
	}
	return nil
}

type genCodesSessionDriver struct{ *gdb.DriverDefault }

var genCodesSessionPool *sql.DB

func (d *genCodesSessionDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &genCodesSessionDriver{&gdb.DriverDefault{Core: core}}, nil
}
func (*genCodesSessionDriver) Master(...string) (*sql.DB, error) { return genCodesSessionPool, nil }
func (*genCodesSessionDriver) Slave(...string) (*sql.DB, error)  { return genCodesSessionPool, nil }
func (*genCodesSessionDriver) TableFields(context.Context, string, ...string) (map[string]*gdb.TableField, error) {
	return map[string]*gdb.TableField{"id": {Name: "id", Type: "bigint", Index: 0}, "status": {Name: "status", Type: "int", Index: 1}}, nil
}

func TestGenCodesStatusSurvivesConnectionReplacement(t *testing.T) {
	const childEnv = "HOTGO_GEN_TRANSACTION_TEST"
	if os.Getenv(childEnv) != "1" {
		cmd := exec.Command(os.Args[0], "-test.run=^TestGenCodesStatusSurvivesConnectionReplacement$", "-test.v")
		cmd.Env = append(os.Environ(), childEnv+"=1")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("transaction regression failed: %v\n%s", err, output)
		}
		return
	}
	t.Chdir(t.TempDir())
	disableGenCodesFileLogs(t)
	const driverName = "gencodes_session_test"
	if err := gdb.Register(driverName, &genCodesSessionDriver{}); err != nil {
		t.Fatal(err)
	}
	if err := gdb.SetConfigGroup("default", gdb.ConfigGroup{{Type: driverName}}); err != nil {
		t.Fatal(err)
	}
	for _, status := range []int{consts.GenCodesStatusOk, consts.GenCodesStatusFail, consts.GenCodesStatusRunning} {
		t.Run(consts.GenCodesStatusNameMap[status], func(t *testing.T) {
			store := &genCodesSessionStore{status: consts.GenCodesStatusRunning}
			genCodesSessionPool = sql.OpenDB(genCodesSessionConnector{store})
			defer genCodesSessionPool.Close()
			genCodesSessionPool.SetMaxOpenConns(1)
			ctx := context.Background()
			if err := NewSysGenCodes().Status(ctx, &sysin.GenCodesStatusInp{Id: 56, Status: status}); err != nil {
				t.Fatalf("status save failed: %+v", err)
			}
			immediate, err := g.DB().GetValue(ctx, "SELECT status FROM hg_sys_gen_codes WHERE id=56")
			if err != nil || immediate.Int() != status {
				t.Fatalf("immediate result=%v err=%v", immediate, err)
			}
			genCodesSessionPool.SetMaxIdleConns(0) // 丢弃旧连接未提交事务，模拟热重启。
			fresh, err := g.DB().GetValue(ctx, "SELECT status FROM hg_sys_gen_codes WHERE id=56")
			if err != nil || fresh.Int() != status {
				t.Fatalf("status reverted after reconnect: want=%d got=%v err=%v", status, fresh, err)
			}
		})
	}
	t.Run("menu publish commits before success and survives reconnect", func(t *testing.T) {
		store := &genCodesSessionStore{status: consts.GenCodesStatusRunning}
		genCodesSessionPool = sql.OpenDB(genCodesSessionConnector{store})
		defer genCodesSessionPool.Close()
		genCodesSessionPool.SetMaxOpenConns(1)
		ctx := context.Background()
		dir := t.TempDir()
		codePath := filepath.Join(dir, "index.vue")
		sqlPath := filepath.Join(dir, "menu.sql")
		service := NewSysGenCodes()
		preview := &sysin.GenCodesPreviewModel{Views: map[string]*sysin.GenFile{
			"web.index.vue": {Path: codePath, Content: "generated", Meth: consts.GenCodesBuildMethCreate},
			"source.sql":    {Path: sqlPath, Content: "SET SQL_MODE = \"NO_AUTO_VALUE_ON_ZERO\";\nSET AUTOCOMMIT = 0;\nSTART TRANSACTION;\nINSERT INTO menus VALUES(NULL);\nCOMMIT;", Meth: consts.GenCodesBuildMethCreate},
		}}
		if err := service.publishBuildPreview(ctx, 56, preview, func(context.Context) error {
			if _, err := os.Stat(codePath); !os.IsNotExist(err) {
				t.Error("code written before DAO step")
			}
			return nil
		}, func(context.Context) error {
			if _, err := os.Stat(codePath); err != nil {
				t.Error("service step ran before code write")
			}
			if _, err := os.Stat(sqlPath); !os.IsNotExist(err) {
				t.Error("menu written before service step")
			}
			return nil
		}); err != nil {
			t.Fatal(err)
		}
		if store.menus != 1 {
			t.Fatalf("menu transaction not committed: %+v", store)
		}
		// 再次覆盖同一份菜单 SQL 时仍应导入菜单，但不重写相同文件。
		oldTime := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
		if err := os.Chtimes(sqlPath, oldTime, oldTime); err != nil {
			t.Fatal(err)
		}
		before, err := os.Stat(sqlPath)
		if err != nil {
			t.Fatal(err)
		}
		noop := func(context.Context) error { return nil }
		if err = service.publishBuildPreview(ctx, 56, preview, noop, noop); err != nil {
			t.Fatal(err)
		}
		after, err := os.Stat(sqlPath)
		if err != nil {
			t.Fatal(err)
		}
		if !before.ModTime().Equal(after.ModTime()) || store.menus != 2 {
			t.Fatalf("unchanged SQL rewritten or menu import skipped: before=%v after=%v menus=%d", before.ModTime(), after.ModTime(), store.menus)
		}
		if err := service.Status(ctx, &sysin.GenCodesStatusInp{Id: 56, Status: consts.GenCodesStatusOk}); err != nil {
			t.Fatal(err)
		}
		genCodesSessionPool.SetMaxIdleConns(0)
		fresh, err := g.DB().GetValue(ctx, "SELECT status FROM hg_sys_gen_codes WHERE id=56")
		if err != nil || fresh.Int() != consts.GenCodesStatusOk {
			t.Fatalf("success lost after menu import/reconnect: %v %v", fresh, err)
		}
		for _, statement := range store.statements {
			if strings.Contains(statement, "AUTOCOMMIT") || strings.Contains(statement, "SQL_MODE") {
				t.Fatalf("session polluted by %s", statement)
			}
		}
	})
	t.Run("publish failure preserves files and rolls back menus", func(t *testing.T) {
		for _, failService := range []bool{false, true} {
			store := &genCodesSessionStore{status: consts.GenCodesStatusRunning}
			genCodesSessionPool = sql.OpenDB(genCodesSessionConnector{store})
			dir := t.TempDir()
			codePath, sqlPath, skippedPath := filepath.Join(dir, "new.vue"), filepath.Join(dir, "menu.sql"), filepath.Join(dir, "manual.vue")
			if err := os.WriteFile(skippedPath, []byte("manual content"), 0644); err != nil {
				t.Fatal(err)
			}
			preview := &sysin.GenCodesPreviewModel{Views: map[string]*sysin.GenFile{
				"web.index.vue": {Path: codePath, Content: "generated", Meth: consts.GenCodesBuildMethCreate},
				"web.edit.vue":  {Path: skippedPath, Content: "wrong overwrite", Meth: consts.GenCodesBuildMethSkip},
				"source.sql":    {Path: sqlPath, Content: "INSERT INTO menus VALUES(NULL);\nCOMMIT;\nFAIL;", Meth: consts.GenCodesBuildMethCreate},
			}}
			err := NewSysGenCodes().publishBuildPreview(context.Background(), 56, preview, func(context.Context) error { return nil }, func(context.Context) error {
				if failService {
					return errors.New("service publish failed")
				}
				return nil
			})
			genCodesSessionPool.Close()
			if err == nil || store.menus != 0 {
				t.Fatalf("partial publish reported success: err=%v menus=%d", err, store.menus)
			}
			if _, err := os.Stat(sqlPath); !os.IsNotExist(err) {
				t.Error("failed menu SQL retained")
			}
			content, err := os.ReadFile(skippedPath)
			if err != nil || string(content) != "manual content" {
				t.Error("existing file overwritten")
			}
		}
	})
	t.Run("commit failure is returned", func(t *testing.T) {
		store := &genCodesSessionStore{status: consts.GenCodesStatusRunning, failCommit: true}
		genCodesSessionPool = sql.OpenDB(genCodesSessionConnector{store})
		defer genCodesSessionPool.Close()
		err := NewSysGenCodes().Status(context.Background(), &sysin.GenCodesStatusInp{Id: 56, Status: consts.GenCodesStatusOk})
		if err == nil || store.status != consts.GenCodesStatusRunning {
			t.Fatalf("uncommitted result reported as success: err=%v persisted=%d", err, store.status)
		}
	})
}
