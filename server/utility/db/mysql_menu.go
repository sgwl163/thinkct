package db

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// ImportMysqlMenuSQL 导入生成器菜单脚本。事务由 GoFrame 管理，不执行脚本中的事务和连接设置。
// 仅用于项目生成的逐行菜单 SQL，不作为任意 SQL 文件的通用导入器。
func ImportMysqlMenuSQL(ctx context.Context, database gdb.DB, content string) error {
	return database.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, line := range strings.Split(content, "\n") {
			statement := strings.TrimSpace(line)
			if statement == "" || strings.HasPrefix(statement, "--") {
				continue
			}
			// 导出脚本可独立执行，但在框架事务内不能再次控制事务或改变连接设置。
			upper := strings.ToUpper(strings.TrimSpace(strings.TrimSuffix(statement, ";")))
			if upper == "START TRANSACTION" || upper == "BEGIN" || upper == "COMMIT" || upper == "ROLLBACK" {
				continue
			}
			setting := strings.ToUpper(strings.Join(strings.Fields(strings.SplitN(statement, "=", 2)[0]), " "))
			switch setting {
			case "SET AUTOCOMMIT", "SET SESSION AUTOCOMMIT", "SET @@SESSION.AUTOCOMMIT",
				"SET SQL_MODE", "SET SESSION SQL_MODE", "SET @@SESSION.SQL_MODE":
				continue
			}
			if _, err := tx.Exec(statement); err != nil {
				return gerror.Wrap(err, "导入生成菜单失败")
			}
		}
		return nil
	})
}
