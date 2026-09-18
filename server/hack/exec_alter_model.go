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
ALTER TABLE hg_openai_api_model
ADD COLUMN is_vision tinyint(1) DEFAULT 0 COMMENT '是否支持识图' AFTER model,
ADD COLUMN is_file tinyint(1) DEFAULT 0 COMMENT '是否支持上传文件' AFTER is_vision,
ADD COLUMN is_search tinyint(1) DEFAULT 0 COMMENT '是否支持联网搜索' AFTER is_file;

UPDATE hg_openai_api_model SET is_vision = 1 WHERE model LIKE '%vision%' OR model LIKE '%gpt-4o%';
`
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SQL executed successfully.")
}