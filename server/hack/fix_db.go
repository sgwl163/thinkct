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

	// Update the latest 4 records with the task ID and mock buttons JSON
	mockButtons := `[{"customId": "MJ::JOB::upsample::1::31307a65-856c-4b73-a621-0a6713833215", "emoji": "", "label": "U1", "type": 2, "style": 2}, {"customId": "MJ::JOB::variation::1::31307a65-856c-4b73-a621-0a6713833215", "emoji": "", "label": "V1", "type": 2, "style": 2}]`
	_, err = db.Exec("UPDATE hg_openai_draw_task SET mj_task_id = ?, buttons = ? WHERE id IN (21, 20, 19, 18)", "1784915543848681", mockButtons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fixed DB records successfully!")
}