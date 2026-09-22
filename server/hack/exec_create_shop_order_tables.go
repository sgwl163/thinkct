package main

import (
	"fmt"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func main() {
	ctx := gctx.New()
	sql := gfile.GetContents("storage/data/generate/shop_order.sql")
	if sql == "" {
		fmt.Println("Error: shop order SQL file is empty or unavailable")
		return
	}

	for _, statement := range strings.Split(sql, ";") {
		if statement = strings.TrimSpace(statement); statement == "" {
			continue
		}
		if _, err := g.DB().Exec(ctx, statement); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	fmt.Println("Successfully created shop order tables.")
}
