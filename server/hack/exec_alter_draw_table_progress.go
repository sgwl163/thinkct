package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	_, err := g.DB().Exec(ctx, "ALTER TABLE `hg_openai_draw_task` ADD COLUMN `progress` int(11) NOT NULL DEFAULT 0 COMMENT '进度(0-100)';")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully added progress column.")
	}
}
