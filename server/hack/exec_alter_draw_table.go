package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/develop_hotgo?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := `
ALTER TABLE hg_openai_draw_task 
ADD COLUMN negative_prompt text COMMENT '负面提示词',
ADD COLUMN stylize int DEFAULT 0 COMMENT '风格化程度(0-1000)',
ADD COLUMN chaos int DEFAULT 0 COMMENT '多样性(0-100)',
ADD COLUMN seed int DEFAULT 0 COMMENT '随机种子',
ADD COLUMN mj_task_id varchar(100) DEFAULT '' COMMENT 'MJ远端任务ID',
ADD COLUMN buttons text COMMENT 'MJ操作按钮JSON';
`
	_, err = db.Exec(queries)
	if err != nil {
		fmt.Println("Warning or Error:", err)
	} else {
		fmt.Println("hg_openai_draw_task table altered successfully!")
	}
}