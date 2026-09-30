package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const userLogAlterSQL = `ALTER TABLE hg_user_log
	CHANGE COLUMN login_at event_at datetime DEFAULT NULL COMMENT '事件时间',
	CHANGE COLUMN login_ip ip varchar(128) NOT NULL DEFAULT '' COMMENT 'IP地址',
	CHANGE COLUMN err_msg message varchar(1000) NOT NULL DEFAULT '' COMMENT '结果消息',
	ADD COLUMN log_type varchar(32) NOT NULL DEFAULT 'login' COMMENT '记录类型:login=登录,operation=操作' AFTER username,
	ADD COLUMN action varchar(64) NOT NULL DEFAULT 'login' COMMENT '操作标识' AFTER log_type,
	ADD COLUMN action_name varchar(128) NOT NULL DEFAULT '用户登录' COMMENT '操作名称' AFTER action,
	ADD COLUMN extra text COMMENT '扩展数据' AFTER message,
	DROP INDEX idx_user_created_at,
	ADD KEY idx_user_type_event_at (user_id, log_type, event_at),
	ADD KEY idx_action_event_at (action, event_at),
	COMMENT='用户_统一行为记录'`

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/develop_hotgo?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var oldTable string
	err = db.QueryRow(`SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'hg_user_login_log'`).Scan(&oldTable)
	if err == sql.ErrNoRows {
		fmt.Println("No legacy hg_user_login_log table found; no migration is needed.")
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	var newTable string
	err = db.QueryRow(`SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'hg_user_log'`).Scan(&newTable)
	if err == nil {
		log.Fatal("hg_user_log already exists; resolve the duplicate table before running this migration")
	}
	if err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if _, err = db.Exec("RENAME TABLE hg_user_login_log TO hg_user_log"); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(userLogAlterSQL); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrated hg_user_login_log to hg_user_log successfully.")
}
