package main

import (
	"flag"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"os/exec"
)

func main() {
	apply := flag.Bool("apply", false, "Create cart schema")
	generate := flag.Bool("generate", false, "Generate only cart DAO using project CLI configuration")
	flag.Parse()
	ctx := gctx.New()
	if *generate {
		fmt.Println("DAO scope preview: gf gen dao --tables hg_shop_cart --removePrefix hg_ --path internal (clear=false; configuration from hack/config.yaml)")
		cmd := exec.Command("gf", "gen", "dao", "--tables", "hg_shop_cart", "--removePrefix", "hg_", "--path", "internal")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		return
	}
	db := g.DB()
	db.SetDebug(false)
	query := `SELECT t.TABLE_NAME,t.ENGINE,c.COLUMN_NAME,c.COLUMN_TYPE,s.INDEX_NAME,s.NON_UNIQUE FROM information_schema.TABLES t LEFT JOIN information_schema.COLUMNS c ON c.TABLE_SCHEMA=t.TABLE_SCHEMA AND c.TABLE_NAME=t.TABLE_NAME LEFT JOIN information_schema.STATISTICS s ON s.TABLE_SCHEMA=c.TABLE_SCHEMA AND s.TABLE_NAME=c.TABLE_NAME AND s.COLUMN_NAME=c.COLUMN_NAME WHERE t.TABLE_SCHEMA=DATABASE() AND (t.TABLE_NAME LIKE '%cart%' OR t.TABLE_NAME IN ('hg_user','hg_shop_product','hg_shop_order','hg_shop_order_item','hg_user_credits_log')) ORDER BY t.TABLE_NAME,c.ORDINAL_POSITION,s.INDEX_NAME`
	rows, err := db.GetAll(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows.Json())
	if !*apply {
		return
	}
	if _, err = db.Exec(ctx, gfile.GetContents("storage/data/generate/shop_cart.sql")); err != nil {
		panic(err)
	}
	fmt.Println("Shop cart schema ready; no business rows read or changed")
}
