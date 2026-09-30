package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type columnMigration struct {
	legacyName string
	targetName string
	definition string
}

type columnInfo struct {
	exists   bool
	dataType string
}

var userDateColumns = []columnMigration{
	{legacyName: "last_login_time", targetName: "last_login_at", definition: "DATETIME DEFAULT NULL COMMENT '上次登录时间'"},
	{legacyName: "join_time", targetName: "joined_at", definition: "DATETIME DEFAULT NULL COMMENT '加入时间'"},
	{legacyName: "update_time", targetName: "updated_at", definition: "DATETIME DEFAULT NULL COMMENT '更新时间'"},
	{legacyName: "create_time", targetName: "created_at", definition: "DATETIME DEFAULT NULL COMMENT '创建时间'"},
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/develop_hotgo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, column := range userDateColumns {
		if err = migrateColumn(db, column); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Migrated hg_user time columns to DATETIME *_at columns.")
}

func migrateColumn(db *sql.DB, column columnMigration) error {
	target, err := getColumnInfo(db, column.targetName)
	if err != nil {
		return err
	}
	legacy, err := getColumnInfo(db, column.legacyName)
	if err != nil {
		return err
	}

	switch {
	case target.exists && target.dataType == "datetime":
		if !legacy.exists {
			return nil
		}
		if err = copyColumnData(db, column.targetName, column.legacyName, legacy.dataType); err != nil {
			return err
		}
		_, err = db.Exec(fmt.Sprintf("ALTER TABLE `hg_user` DROP COLUMN `%s`", column.legacyName))
		return err
	case target.exists && target.dataType == "bigint":
		return convertTargetColumn(db, column)
	case !target.exists && legacy.exists:
		if err = addTargetColumn(db, column, column.legacyName); err != nil {
			return err
		}
		if err = copyColumnData(db, column.targetName, column.legacyName, legacy.dataType); err != nil {
			return err
		}
		_, err = db.Exec(fmt.Sprintf("ALTER TABLE `hg_user` DROP COLUMN `%s`", column.legacyName))
		return err
	case target.exists:
		return fmt.Errorf("hg_user.%s 的类型为 %s，无法自动迁移", column.targetName, target.dataType)
	default:
		return fmt.Errorf("hg_user 缺少字段 %s 和 %s", column.legacyName, column.targetName)
	}
}

func convertTargetColumn(db *sql.DB, column columnMigration) error {
	temporary := column.targetName + "_migrating"
	temporaryInfo, err := getColumnInfo(db, temporary)
	if err != nil {
		return err
	}
	if !temporaryInfo.exists {
		if err = addTargetColumn(db, columnMigration{targetName: temporary, definition: "DATETIME DEFAULT NULL"}, column.targetName); err != nil {
			return err
		}
	}
	if err = copyColumnData(db, temporary, column.targetName, "bigint"); err != nil {
		return err
	}
	_, err = db.Exec(fmt.Sprintf(
		"ALTER TABLE `hg_user` DROP COLUMN `%s`, CHANGE COLUMN `%s` `%s` %s",
		column.targetName,
		temporary,
		column.targetName,
		column.definition,
	))
	return err
}

func addTargetColumn(db *sql.DB, column columnMigration, after string) error {
	_, err := db.Exec(fmt.Sprintf(
		"ALTER TABLE `hg_user` ADD COLUMN `%s` %s AFTER `%s`",
		column.targetName,
		column.definition,
		after,
	))
	return err
}

func copyColumnData(db *sql.DB, target, source, sourceType string) error {
	value := fmt.Sprintf("`%s`", source)
	if sourceType == "bigint" {
		value = fmt.Sprintf("CASE WHEN `%s` > 0 THEN FROM_UNIXTIME(`%s`) END", source, source)
	}
	_, err := db.Exec(fmt.Sprintf(
		"UPDATE `hg_user` SET `%s` = COALESCE(`%s`, %s)",
		target,
		target,
		value,
	))
	return err
}

func getColumnInfo(db *sql.DB, column string) (columnInfo, error) {
	var dataType string
	err := db.QueryRow(`
SELECT data_type
FROM information_schema.columns
WHERE table_schema = DATABASE() AND table_name = 'hg_user' AND column_name = ?`, column).Scan(&dataType)
	if err == sql.ErrNoRows {
		return columnInfo{}, nil
	}
	if err != nil {
		return columnInfo{}, err
	}
	return columnInfo{exists: true, dataType: dataType}, nil
}
