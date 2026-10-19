package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/database/gdb"
)

type menuTestConnector struct{ state *menuTestState }
type menuTestState struct {
	committed, pending int
	autocommit         bool
	statements         []string
}

func (c menuTestConnector) Connect(context.Context) (driver.Conn, error) {
	return &menuTestConn{c.state}, nil
}
func (c menuTestConnector) Driver() driver.Driver { return menuTestSQLDriver{} }

type menuTestSQLDriver struct{}

func (menuTestSQLDriver) Open(string) (driver.Conn, error) { return nil, errors.New("use connector") }

type menuTestConn struct{ state *menuTestState }

func (*menuTestConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("unexpected prepare")
}
func (*menuTestConn) Close() error                { return nil }
func (c *menuTestConn) Begin() (driver.Tx, error) { return c, nil }
func (c *menuTestConn) Commit() error {
	c.state.committed += c.state.pending
	c.state.pending = 0
	return nil
}
func (c *menuTestConn) Rollback() error { c.state.pending = 0; return nil }
func (c *menuTestConn) ExecContext(ctx context.Context, query string, _ []driver.NamedValue) (driver.Result, error) {
	c.state.statements = append(c.state.statements, query)
	upper := strings.ToUpper(query)
	switch {
	case strings.HasPrefix(upper, "SET AUTOCOMMIT"):
		c.state.autocommit = false
	case upper == "COMMIT;":
		c.Commit()
	case strings.HasPrefix(upper, "INSERT "):
		c.state.pending++
	case upper == "FAIL;":
		return nil, errors.New("menu insert rejected")
	}
	return driver.RowsAffected(1), nil
}

type menuTestDriver struct{ *gdb.DriverDefault }

var menuTestPool *sql.DB

func (d *menuTestDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &menuTestDriver{&gdb.DriverDefault{Core: core}}, nil
}
func (*menuTestDriver) Master(...string) (*sql.DB, error) { return menuTestPool, nil }
func TestImportMysqlMenuSQL(t *testing.T) {
	const name = "mysql_menu_test"
	if err := gdb.Register(name, &menuTestDriver{}); err != nil {
		t.Fatal(err)
	}
	for _, failure := range []bool{false, true} {
		t.Run(map[bool]string{false: "success preserves session", true: "failure rolls back all menus"}[failure], func(t *testing.T) {
			state := &menuTestState{autocommit: true}
			menuTestPool = sql.OpenDB(menuTestConnector{state})
			defer menuTestPool.Close()
			database, err := gdb.New(gdb.ConfigNode{Type: name})
			if err != nil {
				t.Fatal(err)
			}
			content := "-- generated menu\nSET SQL_MODE = \"NO_AUTO_VALUE_ON_ZERO\";\nSET AUTOCOMMIT = 0;\nSTART TRANSACTION;\nSET @now := now();\nINSERT INTO menus VALUES(NULL,@now);\nSET @dirId = LAST_INSERT_ID();\nINSERT INTO menus VALUES(NULL,@dirId);\nCOMMIT;\n"
			if failure {
				content += "FAIL;\n"
			}
			err = ImportMysqlMenuSQL(context.Background(), database, content)
			if failure && err == nil {
				t.Error("import error swallowed")
			}
			if !failure && err != nil {
				t.Fatal(err)
			}
			if !state.autocommit {
				t.Error("menu import leaked AUTOCOMMIT=0 into pooled connection")
			}
			want := 2
			if failure {
				want = 0
			}
			if state.committed != want {
				t.Errorf("want committed menus=%d got=%d", want, state.committed)
			}
			if !strings.Contains(strings.Join(state.statements, "\n"), "SET @dirId = LAST_INSERT_ID();") {
				t.Error("menu parent variable was removed")
			}
		})
	}
}
