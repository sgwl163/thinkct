package main

import (
	"context"
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
	carouselTable   = "hg_web_carousel"
	carouselSQLFile = "storage/data/generate/web_carousel.sql"
)

// Default: read-only inspection. -apply: create hg_web_carousel if absent and
// seed only while the table is empty. -generate: scoped DAO generation only.
func main() {
	apply := flag.Bool("apply", false, "Create hg_web_carousel if absent and seed it only when the table is empty")
	generate := flag.Bool("generate", false, "Generate only the hg_web_carousel DAO using hack/config.yaml")
	flag.Parse()
	if *generate {
		if *apply {
			carouselFail("use -generate separately from -apply")
		}
		fmt.Println("gf gen dao --tables hg_web_carousel --removePrefix hg_ --path internal (hack/config.yaml; clear=false)")
		cmd := exec.Command("gf", "gen", "dao", "--tables", carouselTable, "--removePrefix", "hg_", "--path", "internal")
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			carouselFail("scoped DAO generation failed: " + err.Error())
		}
		return
	}
	ctx := gctx.New()
	db := g.DB()
	db.SetDebug(false)
	exists, err := db.GetValue(ctx, "SELECT COUNT(*) FROM information_schema.TABLES WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=?", carouselTable)
	carouselCheck(err)
	fmt.Printf("table_state %s exists=%v\n", carouselTable, exists.Int() == 1)
	if exists.Int() == 1 {
		carouselReport(ctx, db)
	}
	if !*apply {
		fmt.Println("inspect only; no schema or data changes")
		return
	}
	statements := carouselStatements()
	if _, err = db.Exec(ctx, statements[0]); err != nil {
		carouselFail("create table failed: " + err.Error())
	}
	result, err := db.Exec(ctx, statements[1])
	if err != nil {
		carouselFail("seed failed: " + err.Error())
	}
	affected, err := result.RowsAffected()
	carouselCheck(err)
	if affected == 0 {
		fmt.Println("seed skipped: table already contains rows; existing data untouched")
	} else {
		fmt.Printf("seed inserted rows=%d\n", affected)
	}
	carouselReport(ctx, db)
}

// carouselStatements accepts only the guarded CREATE + empty-table INSERT for hg_web_carousel.
func carouselStatements() []string {
	var statements []string
	for _, part := range strings.Split(gfile.GetContents(carouselSQLFile), ";") {
		var lines []string
		for _, line := range strings.Split(part, "\n") {
			if trimmed := strings.TrimSpace(line); trimmed != "" && !strings.HasPrefix(trimmed, "--") {
				lines = append(lines, strings.TrimRight(line, "\r"))
			}
		}
		if len(lines) > 0 {
			statements = append(statements, strings.Join(lines, "\n"))
		}
	}
	if len(statements) != 2 ||
		!strings.HasPrefix(statements[0], "CREATE TABLE IF NOT EXISTS `"+carouselTable+"`") ||
		!strings.HasPrefix(statements[1], "INSERT INTO `"+carouselTable+"`") ||
		!strings.Contains(statements[1], "WHERE NOT EXISTS (SELECT 1 FROM `"+carouselTable+"`") {
		carouselFail("unexpected statements in " + carouselSQLFile + "; refusing to execute")
	}
	return statements
}

func carouselReport(ctx context.Context, db gdb.DB) {
	columns, err := db.GetAll(ctx, "SELECT COLUMN_NAME,COLUMN_TYPE,IS_NULLABLE,COLUMN_DEFAULT,EXTRA FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=? ORDER BY ORDINAL_POSITION", carouselTable)
	carouselCheck(err)
	fmt.Println("columns", columns.Json())
	rows, err := db.GetAll(ctx, "SELECT id,title,cta_text,cta_link,color,image,alt,aspect_ratio,sort,status,deleted_at FROM `hg_web_carousel` ORDER BY sort,id")
	carouselCheck(err)
	fmt.Printf("rows=%d %s\n", len(rows), rows.Json())
}

func carouselCheck(err error) {
	if err != nil {
		carouselFail(err.Error())
	}
}

func carouselFail(message string) { fmt.Fprintln(os.Stderr, message); os.Exit(1) }
