package main

import (
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	result, err := g.DB().Model("hg_admin_menu").
		Ctx(ctx).
		Where("name", "reportIndex").
		Data(g.Map{"component": "/forum/report/index"}).
		Update()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Successfully updated %d forum report menu record(s).\n", rowsAffected)
}
