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
CREATE TABLE IF NOT EXISTS hg_openai_music_task (
  id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  user_id bigint(20) NOT NULL COMMENT '用户ID',
  prompt text NOT NULL COMMENT '提示词(灵感或歌词)',
  type tinyint(1) DEFAULT '1' COMMENT '创作模式: 1=灵感模式, 2=定制模式',
  make_instrumental tinyint(1) DEFAULT '0' COMMENT '纯音乐: 0=否 1=是',
  title varchar(255) DEFAULT '' COMMENT '歌名',
  tags varchar(255) DEFAULT '' COMMENT '风格标签',
  model varchar(255) NOT NULL COMMENT '使用的模型(mv)',
  audio_url text COMMENT '生成的音频URL',
  image_url text COMMENT '封面图URL',
  video_url text COMMENT '生成的视频URL',
  duration float DEFAULT '0' COMMENT '时长',
  status tinyint(1) DEFAULT '0' COMMENT '状态: 0=进行中 1=成功, 2=失败', 
  fail_reason text COMMENT '失败原因',
  remote_task_id varchar(100) DEFAULT '' COMMENT '远端任务ID',
  remote_clip_id varchar(100) DEFAULT '' COMMENT '远端歌曲ID',
  created_at datetime DEFAULT NULL COMMENT '创建时间',
  updated_at datetime DEFAULT NULL COMMENT '修改时间',
  deleted_at datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  KEY idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI音乐任务';
`
	_, err := g.DB().Exec(ctx, queries)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("hg_openai_music_task table created successfully!")
	}
}
