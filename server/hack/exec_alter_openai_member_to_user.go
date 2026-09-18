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
-- 1. 更新 hg_openai_chat_folder 表
ALTER TABLE hg_openai_chat_folder 
CHANGE COLUMN member_id user_id bigint(20) NOT NULL COMMENT '用户ID';
ALTER TABLE hg_openai_chat_folder 
DROP INDEX idx_member_id,
ADD INDEX idx_user_id (user_id);

-- 2. 更新 hg_openai_chat_session 表
ALTER TABLE hg_openai_chat_session 
CHANGE COLUMN member_id user_id bigint(20) NOT NULL COMMENT '用户ID';
ALTER TABLE hg_openai_chat_session 
DROP INDEX idx_member_id,
ADD INDEX idx_user_id (user_id);

-- 3. 更新 hg_openai_chat_message 表
ALTER TABLE hg_openai_chat_message 
CHANGE COLUMN member_id user_id bigint(20) NOT NULL COMMENT '用户ID';
ALTER TABLE hg_openai_chat_message 
DROP INDEX idx_member_id,
ADD INDEX idx_user_id (user_id);
`
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OpenAI tables updated successfully!")
}
