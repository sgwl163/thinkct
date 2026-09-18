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
ALTER TABLE hg_openai_chat_session
ADD COLUMN is_top tinyint(1) DEFAULT 0 COMMENT '是否置顶' AFTER title;
`
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SQL executed successfully.")
}
