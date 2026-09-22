package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type permissionMenu struct {
	Title       string
	Name        string
	Permissions string
	Sort        int
}

func main() {
	ctx := gctx.New()
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, menu := range []permissionMenu{
			{Title: "删除生成配置", Name: "genCodesDelete", Permissions: "/genCodes/delete", Sort: 10},
			{Title: "删除生成代码", Name: "genCodesDeleteBuild", Permissions: "/genCodes/deleteBuild", Sort: 20},
		} {
			count, err := tx.Model("hg_admin_menu").Where("permissions", menu.Permissions).Count()
			if err != nil {
				return err
			}
			if count > 0 {
				continue
			}
			if _, err = tx.Model("hg_admin_menu").Data(g.Map{
				"pid":         2098,
				"level":       3,
				"tree":        "tr_2097 tr_2098 ",
				"title":       menu.Title,
				"name":        menu.Name,
				"type":        3,
				"permissions": menu.Permissions,
				"always_show": 1,
				"is_frame":    2,
				"sort":        menu.Sort,
				"status":      1,
			}).Insert(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully added code generation delete permissions.")
}
