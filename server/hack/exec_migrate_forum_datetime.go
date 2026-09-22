package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type timeColumn struct {
	Source     string
	Target     string
	Definition string
}

type indexMigration struct {
	Source     string
	Target     string
	Definition string
}

type tableMigration struct {
	Name       string
	Columns    []timeColumn
	AddColumns []columnMigration
	Indexes    []indexMigration
}

type columnMigration struct {
	Name       string
	Definition string
	After      string
}

var forumMigrations = []tableMigration{
	{
		Name: "hg_forum_category",
		Columns: []timeColumn{
			{Source: "update_time", Target: "updated_at", Definition: "DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'"},
			{Source: "create_time", Target: "created_at", Definition: "DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'"},
		},
	},
	{
		Name: "hg_forum_topic",
		Columns: []timeColumn{
			{Source: "last_reply_time", Target: "last_reply_at", Definition: "DATETIME DEFAULT NULL COMMENT '最后回复时间'"},
			{Source: "update_time", Target: "updated_at", Definition: "DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'"},
			{Source: "create_time", Target: "created_at", Definition: "DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'"},
			{Source: "delete_time", Target: "deleted_at", Definition: "DATETIME DEFAULT NULL COMMENT '删除时间'"},
		},
		AddColumns: []columnMigration{
			{Name: "keywords", Definition: "VARCHAR(500) NOT NULL DEFAULT '' COMMENT '主题关键词，英文逗号分隔'", After: "title"},
			{Name: "description", Definition: "VARCHAR(500) NOT NULL DEFAULT '' COMMENT '主题摘要'", After: "keywords"},
		},
		Indexes: []indexMigration{
			{Source: "idx_category_status_top_time", Target: "idx_category_status_top_created_at", Definition: "(category_id, status, is_top, created_at)"},
			{Source: "idx_user_create_time", Target: "idx_user_created_at", Definition: "(user_id, created_at)"},
			{Source: "idx_status_reply_time", Target: "idx_status_last_reply_at", Definition: "(status, last_reply_at)"},
		},
	},
	{
		Name: "hg_forum_comment",
		Columns: []timeColumn{
			{Source: "update_time", Target: "updated_at", Definition: "DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'"},
			{Source: "create_time", Target: "created_at", Definition: "DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'"},
			{Source: "delete_time", Target: "deleted_at", Definition: "DATETIME DEFAULT NULL COMMENT '删除时间'"},
		},
		Indexes: []indexMigration{
			{Source: "idx_topic_status_time", Target: "idx_topic_status_created_at", Definition: "(topic_id, status, created_at)"},
			{Source: "idx_parent_time", Target: "idx_parent_created_at", Definition: "(parent_id, created_at)"},
			{Source: "idx_user_create_time", Target: "idx_user_created_at", Definition: "(user_id, created_at)"},
		},
	},
	{
		Name: "hg_forum_report",
		Columns: []timeColumn{
			{Source: "handled_time", Target: "handled_at", Definition: "DATETIME DEFAULT NULL COMMENT '处理时间'"},
			{Source: "create_time", Target: "created_at", Definition: "DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'"},
		},
		Indexes: []indexMigration{
			{Source: "idx_user_create_time", Target: "idx_user_created_at", Definition: "(user_id, created_at)"},
			{Source: "idx_status_create_time", Target: "idx_status_created_at", Definition: "(status, created_at)"},
		},
	},
}

func main() {
	ctx := gctx.New()
	for _, migration := range forumMigrations {
		if err := migrateTable(ctx, migration); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	if err := migrateForumInteractions(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := migrateForumUserFollows(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully migrated forum schema to the *_at convention.")
}

func migrateTable(ctx context.Context, migration tableMigration) error {
	for _, column := range migration.Columns {
		exists, err := columnExists(ctx, migration.Name, column.Source)
		if err != nil {
			return err
		}
		if !exists {
			continue
		}
		targetExists, err := columnExists(ctx, migration.Name, column.Target)
		if err != nil {
			return err
		}
		if targetExists {
			return fmt.Errorf("社区表 %s 同时存在 %s 和 %s，请先人工确认数据后再迁移", migration.Name, column.Source, column.Target)
		}
		dataType, err := columnDataType(ctx, migration.Name, column.Source)
		if err != nil {
			return err
		}
		if dataType != "datetime" && dataType != "bigint" {
			return fmt.Errorf("社区表 %s.%s 的时间字段类型为 %s，无法自动迁移", migration.Name, column.Source, dataType)
		}
	}

	if err := dropLegacyIndexes(ctx, migration); err != nil {
		return err
	}
	for _, column := range migration.Columns {
		exists, err := columnExists(ctx, migration.Name, column.Source)
		if err != nil {
			return err
		}
		if !exists {
			continue
		}
		dataType, err := columnDataType(ctx, migration.Name, column.Source)
		if err != nil {
			return err
		}
		if err = migrateColumn(ctx, migration.Name, column, dataType); err != nil {
			return err
		}
	}
	if err := migrateIndexes(ctx, migration); err != nil {
		return err
	}
	return addMissingColumns(ctx, migration)
}

func migrateColumn(ctx context.Context, table string, column timeColumn, dataType string) error {
	if dataType == "datetime" {
		_, err := g.DB().Exec(ctx, fmt.Sprintf("ALTER TABLE `%s` CHANGE COLUMN `%s` `%s` %s", table, column.Source, column.Target, column.Definition))
		return err
	}
	if dataType != "bigint" {
		return fmt.Errorf("社区表 %s.%s 的时间字段类型为 %s，无法自动迁移", table, column.Source, dataType)
	}

	temporary := column.Target + "_migrating"
	statements := []string{
		fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` DATETIME NULL AFTER `%s`", table, temporary, column.Source),
		fmt.Sprintf("UPDATE `%s` SET `%s` = CASE WHEN `%s` > 0 THEN FROM_UNIXTIME(`%s`) END", table, temporary, column.Source, column.Source),
		fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN `%s`, CHANGE COLUMN `%s` `%s` %s", table, column.Source, temporary, column.Target, column.Definition),
	}
	for _, statement := range statements {
		if _, err := g.DB().Exec(ctx, statement); err != nil {
			return err
		}
	}
	return nil
}

func dropLegacyIndexes(ctx context.Context, migration tableMigration) error {
	for _, index := range migration.Indexes {
		exists, err := indexExists(ctx, migration.Name, index.Source)
		if err != nil {
			return err
		}
		if exists {
			if _, err = g.DB().Exec(ctx, fmt.Sprintf("ALTER TABLE `%s` DROP INDEX `%s`", migration.Name, index.Source)); err != nil {
				return err
			}
		}
	}
	return nil
}

func migrateIndexes(ctx context.Context, migration tableMigration) error {
	for _, index := range migration.Indexes {
		targetExists, err := indexExists(ctx, migration.Name, index.Target)
		if err != nil {
			return err
		}
		if !targetExists {
			if _, err = g.DB().Exec(ctx, fmt.Sprintf("ALTER TABLE `%s` ADD KEY `%s` %s", migration.Name, index.Target, index.Definition)); err != nil {
				return err
			}
		}
	}
	return nil
}

func addMissingColumns(ctx context.Context, migration tableMigration) error {
	for _, column := range migration.AddColumns {
		exists, err := columnExists(ctx, migration.Name, column.Name)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		statement := fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s", migration.Name, column.Name, column.Definition)
		if column.After != "" {
			statement += fmt.Sprintf(" AFTER `%s`", column.After)
		}
		if _, err = g.DB().Exec(ctx, statement); err != nil {
			return err
		}
	}
	return nil
}

func migrateForumInteractions(ctx context.Context) error {
	if _, err := g.DB().Exec(ctx, `
CREATE TABLE IF NOT EXISTS hg_forum_interaction (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '互动ID',
    target_type VARCHAR(32) NOT NULL COMMENT '互动对象类型:topic,comment',
    target_id BIGINT UNSIGNED NOT NULL COMMENT '互动对象ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    action_type VARCHAR(32) NOT NULL COMMENT '互动类型:like,favorite',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_target_user_action (target_type, target_id, user_id, action_type),
    KEY idx_target_action (target_type, target_id, action_type),
    KEY idx_user_created_at (user_id, created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛互动表'`); err != nil {
		return err
	}
	if err := migrateTopicInteractionTable(ctx, "hg_forum_topic_like", "like"); err != nil {
		return err
	}
	return migrateTopicInteractionTable(ctx, "hg_forum_topic_favorite", "favorite")
}

func migrateTopicInteractionTable(ctx context.Context, table, actionType string) error {
	exists, err := tableExists(ctx, table)
	if err != nil || !exists {
		return err
	}
	createdAt := "created_at"
	createdAtExists, err := columnExists(ctx, table, createdAt)
	if err != nil {
		return err
	}
	if !createdAtExists {
		createdAt = "create_time"
	}
	statement := fmt.Sprintf(`
INSERT IGNORE INTO hg_forum_interaction (target_type, target_id, user_id, action_type, created_at)
SELECT 'topic', topic_id, user_id, '%s', %s FROM %s`, actionType, createdAt, table)
	if _, err = g.DB().Exec(ctx, statement); err != nil {
		return err
	}
	_, err = g.DB().Exec(ctx, fmt.Sprintf("DROP TABLE `%s`", table))
	return err
}

func migrateForumUserFollows(ctx context.Context) error {
	_, err := g.DB().Exec(ctx, `
CREATE TABLE IF NOT EXISTS hg_forum_user_follow (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关注ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '关注用户ID',
    follow_user_id BIGINT UNSIGNED NOT NULL COMMENT '被关注用户ID',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '关注时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_user_follow_user (user_id, follow_user_id),
    KEY idx_follow_user_created_at (follow_user_id, created_at),
    KEY idx_user_created_at (user_id, created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛用户关注表'`)
	return err
}

func tableExists(ctx context.Context, table string) (bool, error) {
	value, err := g.DB().GetValue(ctx, `
SELECT COUNT(*)
FROM information_schema.tables
WHERE table_schema = DATABASE() AND table_name = ?`, table)
	return value.Int() > 0, err
}

func columnExists(ctx context.Context, table, column string) (bool, error) {
	value, err := g.DB().GetValue(ctx, `
SELECT COUNT(*)
FROM information_schema.columns
WHERE table_schema = DATABASE() AND table_name = ? AND column_name = ?`, table, column)
	return value.Int() > 0, err
}

func columnDataType(ctx context.Context, table, column string) (string, error) {
	value, err := g.DB().GetValue(ctx, `
SELECT data_type
FROM information_schema.columns
WHERE table_schema = DATABASE() AND table_name = ? AND column_name = ?`, table, column)
	return value.String(), err
}

func indexExists(ctx context.Context, table, index string) (bool, error) {
	value, err := g.DB().GetValue(ctx, `
SELECT COUNT(*)
FROM information_schema.statistics
WHERE table_schema = DATABASE() AND table_name = ? AND index_name = ?`, table, index)
	return value.Int() > 0, err
}
