package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	queries := `
CREATE TABLE IF NOT EXISTS hg_openai_video_task (
  id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  user_id bigint(20) NOT NULL COMMENT '用户ID',
  prompt text NOT NULL COMMENT '视频提示词',
  reference_image varchar(1000) DEFAULT '' COMMENT '参考图/首帧URL',
  last_frame varchar(1000) DEFAULT '' COMMENT '尾帧图片URL',
  voice_dub_url varchar(1000) DEFAULT '' COMMENT '自定义配音URL（音频）',
  bg_music_url varchar(1000) DEFAULT '' COMMENT '自定义背景音乐URL（音频）',
  generate_audio tinyint(1) DEFAULT '1' COMMENT '是否生成同步音频: 1=是 0=否',
  model varchar(255) NOT NULL COMMENT '使用的模型',
  video_url text COMMENT '生成的视频URL',
  status tinyint(1) DEFAULT '0' COMMENT '状态: 0=进行中 1=成功, 2=失败', 
  fail_reason text COMMENT '失败原因',
  remote_task_id varchar(100) DEFAULT '' COMMENT '远端任务ID',
  created_at datetime DEFAULT NULL COMMENT '创建时间',
  updated_at datetime DEFAULT NULL COMMENT '修改时间',
  deleted_at datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  KEY idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI视频任务';
`
	_, err := g.DB().Exec(ctx, queries)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("hg_openai_video_task table created successfully!")
	}
}
