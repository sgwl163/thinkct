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
CREATE TABLE IF NOT EXISTS hg_openai_draw_task (
  id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  user_id bigint(20) NOT NULL COMMENT '用户ID',
  prompt text NOT NULL COMMENT '绘画提示词',
  reference_image varchar(1000) DEFAULT '' COMMENT '参考图URL',
  model varchar(255) NOT NULL COMMENT '使用的模型',
  version varchar(50) DEFAULT '' COMMENT '模型版本',
  ratio varchar(50) DEFAULT '1:1' COMMENT '图片比例',
  raw tinyint(1) DEFAULT '0' COMMENT '是否开启RAW模式',
  image_url text COMMENT '生成的图片URL',
  status tinyint(1) DEFAULT '0' COMMENT '状态: 0=进行中, 1=成功, 2=失败',
  fail_reason text COMMENT '失败原因',
  created_at datetime DEFAULT NULL COMMENT '创建时间',
  updated_at datetime DEFAULT NULL COMMENT '修改时间',
  deleted_at datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  KEY idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI绘画任务';
`
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hg_openai_draw_task table created successfully!")
}
