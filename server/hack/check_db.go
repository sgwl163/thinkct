package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/develop_hotgo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, mj_task_id, status FROM hg_openai_draw_task ORDER BY id DESC LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("ID\tMJ_TASK_ID\tSTATUS")
	for rows.Next() {
		var id int
		var mjTaskId, status string
		if err := rows.Scan(&id, &mjTaskId, &status); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%s\t%s\n", id, mjTaskId, status)
	}
}