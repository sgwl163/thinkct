package main

import (
	"flag"
	"fmt"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func main() {
	apply := flag.Bool("apply", false, "Create the frontend recharge table if absent")
	inspectTable := flag.String("table", "", "Inspect only this recharge-related table")
	flag.Parse()
	ctx := gctx.New()
	tables, err := g.DB().GetAll(ctx, "SELECT TABLE_NAME, TABLE_COMMENT, ENGINE FROM information_schema.TABLES WHERE TABLE_SCHEMA = DATABASE() AND (TABLE_NAME LIKE '%recharge%' OR TABLE_NAME IN ('hg_admin_order', 'hg_pay_log', 'hg_user', 'hg_user_credits_log'))")
	if err != nil {
		panic(err)
	}
	fmt.Println(tables.Json())
	for _, table := range tables {
		name := table["TABLE_NAME"].String()
		if *inspectTable != "" && name != *inspectTable {
			continue
		}
		columns, err := g.DB().GetAll(ctx, "SELECT COLUMN_NAME, COLUMN_TYPE, COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ? ORDER BY ORDINAL_POSITION", name)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, columns.Json())
		indexes, err := g.DB().GetAll(ctx, "SELECT INDEX_NAME, NON_UNIQUE, SEQ_IN_INDEX, COLUMN_NAME FROM information_schema.STATISTICS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ? ORDER BY INDEX_NAME, SEQ_IN_INDEX", name)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, "indexes", indexes.Json())
	}
	if !*apply {
		return
	}
	sql := gfile.GetContents("storage/data/generate/user_recharge.sql")
	if strings.TrimSpace(sql) == "" {
		panic("recharge SQL is unavailable")
	}
	for _, statement := range strings.Split(sql, ";") {
		if statement = strings.TrimSpace(statement); statement != "" {
			if _, err := g.DB().Exec(ctx, statement); err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("Frontend recharge schema ready")
}
