package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/develop_hotgo?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlBytes, err := ioutil.ReadFile("storage/data/generate/openai_app.sql")
	if err != nil {
		log.Fatal(err)
	}

	queries := string(sqlBytes)
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SQL executed successfully.")
}
