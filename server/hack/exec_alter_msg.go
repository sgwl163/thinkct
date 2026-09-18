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
ALTER TABLE hg_openai_chat_message
ADD COLUMN feedback tinyint(1) DEFAULT 0 COMMENT '反馈: 0=无, 1=喜欢, 2=不喜欢' AFTER model;
`
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SQL executed successfully.")
}
