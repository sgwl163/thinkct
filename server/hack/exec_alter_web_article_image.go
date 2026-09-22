package main

import (
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	_, err := g.DB().Exec(ctx, "ALTER TABLE `hg_web_article` ADD COLUMN `image` varchar(255) NOT NULL DEFAULT '' COMMENT '图片' AFTER `flag`")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully added image column.")
}
