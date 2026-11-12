package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

const (
	navMenuTable   = "hg_web_nav_menu"
	navMenuSQLFile = "storage/data/generate/web_nav_menu.sql"
)

var navMenuColumns = []string{"id", "pid", "level", "tree", "name", "url", "image", "image_alt", "sort", "status", "created_at", "updated_at", "deleted_at"}

// Re-runnable: creates hg_web_nav_menu when absent and seeds the NavMenu.ts rows
// only while the table is completely empty (soft-deleted rows included).
// Existing rows are never updated or deleted; no other table is touched.
func main() {
	apply := flag.Bool("apply", false, "Create hg_web_nav_menu if absent and seed it only when the table is empty")
	generate := flag.Bool("generate", false, "Generate only the hg_web_nav_menu DAO using hack/config.yaml")
	flag.Parse()
	if *generate {
		if *apply {
			navFail("use -generate separately from -apply")
		}
		fmt.Println("gf gen dao --tables hg_web_nav_menu --removePrefix hg_ --path internal (hack/config.yaml; clear=false)")
		cmd := exec.Command("gf", "gen", "dao", "--tables", navMenuTable, "--removePrefix", "hg_", "--path", "internal")
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			navFail("scoped DAO generation failed: " + err.Error())
		}
		return
	}
	ctx := gctx.New()
	db := g.DB()
	db.SetDebug(false)
	exists := navTableExists(ctx, db)
	fmt.Printf("table_state %s exists=%v\n", navMenuTable, exists)
	if exists {
		navReport(ctx, db)
	}
	if !*apply {
		fmt.Println("inspect only; no schema or data changes")
		return
	}
	createSQL, seedSQL := navLoadStatements()
	if _, err := db.Exec(ctx, createSQL); err != nil {
		navFail("create table failed: " + err.Error())
	}
	navCheckColumns(ctx, db)
	result, err := db.Exec(ctx, seedSQL)
	if err != nil {
		navFail("seed failed: " + err.Error())
	}
	affected, err := result.RowsAffected()
	navCheck(err)
	if affected == 0 {
		fmt.Println("seed skipped: table already contains rows; existing data untouched")
	} else {
		fmt.Printf("seed inserted rows=%d\n", affected)
	}
	navReport(ctx, db)
}

// navLoadStatements reads the SQL file and only accepts exactly the expected
// CREATE + guarded INSERT statements targeting hg_web_nav_menu.
func navLoadStatements() (createSQL, seedSQL string) {
	content := gfile.GetContents(navMenuSQLFile)
	if strings.TrimSpace(content) == "" {
		navFail("navigation SQL is unavailable: " + navMenuSQLFile)
	}
	var statements []string
	for _, part := range strings.Split(content, ";") {
		var lines []string
		for _, line := range strings.Split(part, "\n") {
			line = strings.TrimRight(line, "\r")
			if trimmed := strings.TrimSpace(line); trimmed != "" && !strings.HasPrefix(trimmed, "--") {
				lines = append(lines, line)
			}
		}
		if len(lines) > 0 {
			statements = append(statements, strings.Join(lines, "\n"))
		}
	}
	if len(statements) != 2 ||
		!strings.HasPrefix(statements[0], "CREATE TABLE IF NOT EXISTS `"+navMenuTable+"`") ||
		!strings.HasPrefix(statements[1], "INSERT INTO `"+navMenuTable+"`") ||
		!strings.Contains(statements[1], "WHERE NOT EXISTS (SELECT 1 FROM `"+navMenuTable+"`") {
		navFail("unexpected statements in " + navMenuSQLFile + "; refusing to execute")
	}
	return statements[0], statements[1]
}

func navTableExists(ctx context.Context, db gdb.DB) bool {
	v, err := db.GetValue(ctx, "SELECT COUNT(*) FROM information_schema.TABLES WHERE TABLE_SCHEMA=DATABASE() AND TABLE_TYPE='BASE TABLE' AND TABLE_NAME=?", navMenuTable)
	navCheck(err)
	return v.Int() == 1
}

func navCheckColumns(ctx context.Context, db gdb.DB) {
	values, err := db.GetArray(ctx, "SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=? ORDER BY ORDINAL_POSITION", navMenuTable)
	navCheck(err)
	actual := make([]string, 0, len(values))
	for _, v := range values {
		actual = append(actual, v.String())
	}
	if strings.Join(actual, ",") != strings.Join(navMenuColumns, ",") {
		navFail(fmt.Sprintf("existing %s columns %v differ from expected %v; refusing to seed", navMenuTable, actual, navMenuColumns))
	}
}

func navReport(ctx context.Context, db gdb.DB) {
	columns, err := db.GetAll(ctx, "SELECT COLUMN_NAME,COLUMN_TYPE,IS_NULLABLE,COLUMN_DEFAULT,EXTRA FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=? ORDER BY ORDINAL_POSITION", navMenuTable)
	navCheck(err)
	fmt.Println("columns", columns.Json())
	indexes, err := db.GetAll(ctx, "SELECT INDEX_NAME,NON_UNIQUE,SEQ_IN_INDEX,COLUMN_NAME FROM information_schema.STATISTICS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=? ORDER BY INDEX_NAME,SEQ_IN_INDEX", navMenuTable)
	navCheck(err)
	fmt.Println("indexes", indexes.Json())
	checks := []struct{ label, query string }{
		{"rows_total", "SELECT COUNT(*) FROM `hg_web_nav_menu`"},
		{"rows_active", "SELECT COUNT(*) FROM `hg_web_nav_menu` WHERE deleted_at IS NULL"},
		{"rows_status_1", "SELECT COUNT(*) FROM `hg_web_nav_menu` WHERE status=1"},
		{"roots", "SELECT COUNT(*) FROM `hg_web_nav_menu` WHERE pid=0"},
		{"min_id", "SELECT COALESCE(MIN(id),0) FROM `hg_web_nav_menu`"},
		{"max_id", "SELECT COALESCE(MAX(id),0) FROM `hg_web_nav_menu`"},
		{"distinct_id", "SELECT COUNT(DISTINCT id) FROM `hg_web_nav_menu`"},
		{"auto_increment", "SELECT AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME='hg_web_nav_menu'"},
		{"orphan_children", "SELECT COUNT(*) FROM `hg_web_nav_menu` c LEFT JOIN `hg_web_nav_menu` p ON p.id=c.pid WHERE c.pid<>0 AND p.id IS NULL"},
		{"self_parent", "SELECT COUNT(*) FROM `hg_web_nav_menu` WHERE pid=id"},
		{"root_level_or_tree_mismatch", "SELECT COUNT(*) FROM `hg_web_nav_menu` WHERE pid=0 AND (level<>1 OR LENGTH(tree)<>0)"},
		{"child_level_mismatch", "SELECT COUNT(*) FROM `hg_web_nav_menu` c JOIN `hg_web_nav_menu` p ON p.id=c.pid WHERE c.level<>p.level+1"},
		{"child_tree_mismatch", "SELECT COUNT(*) FROM `hg_web_nav_menu` c JOIN `hg_web_nav_menu` p ON p.id=c.pid WHERE c.tree<>CONCAT(p.tree,'tr_',c.pid,' ') OR LENGTH(c.tree)<>LENGTH(CONCAT(p.tree,'tr_',c.pid,' '))"},
		{"duplicate_sibling_sort", "SELECT COUNT(*) FROM (SELECT pid,sort FROM `hg_web_nav_menu` GROUP BY pid,sort HAVING COUNT(*)>1) d"},
	}
	for _, c := range checks {
		v, err := db.GetValue(ctx, c.query)
		navCheck(err)
		fmt.Printf("%s=%s\n", c.label, v.String())
	}
	levels, err := db.GetAll(ctx, "SELECT level,COUNT(*) AS cnt FROM `hg_web_nav_menu` GROUP BY level ORDER BY level")
	navCheck(err)
	fmt.Println("rows_by_level", levels.Json())
	roots, err := db.GetAll(ctx, "SELECT r.id,r.name,r.sort,(SELECT COUNT(*) FROM `hg_web_nav_menu` x WHERE x.tree LIKE CONCAT('tr_',r.id,' %')) AS descendants FROM `hg_web_nav_menu` r WHERE r.pid=0 ORDER BY r.sort,r.id")
	navCheck(err)
	fmt.Println("roots_detail", roots.Json())
	all, err := db.GetAll(ctx, "SELECT id,pid,level,tree,name,url,image,image_alt,sort FROM `hg_web_nav_menu` ORDER BY id")
	navCheck(err)
	h := sha256.New()
	for _, r := range all {
		_, _ = fmt.Fprintf(h, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", r["id"], r["pid"], r["level"], r["tree"], r["name"], r["url"], r["image"], r["image_alt"], r["sort"])
	}
	fmt.Printf("content_sha256=%x (id,pid,level,tree,name,url,image,image_alt,sort ordered by id)\n", h.Sum(nil))
}

func navCheck(err error) {
	if err != nil {
		navFail(err.Error())
	}
}

func navFail(message string) { fmt.Fprintln(os.Stderr, message); os.Exit(1) }
