package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/frame/g"
)

const oldProductTable = "hg_web_products"
const newProductTable = "hg_shop_product"

// This migration deliberately uses database/sql: no ORM hooks, business writes,
// SQL debug logs or connection strings are emitted.
func main() {
	apply := flag.Bool("apply", false, "Explicitly rename the product table and update exact metadata references")
	generate := flag.Bool("generate", false, "Generate only the new product DAO using hack/config.yaml")
	flag.Parse()
	if *generate {
		if *apply {
			fail("use -generate separately from -apply")
		}
		fmt.Println("gf gen dao --tables hg_shop_product --removePrefix hg_ --path internal (hack/config.yaml; clear=false)")
		cmd := exec.Command("gf", "gen", "dao", "--tables", newProductTable, "--removePrefix", "hg_", "--path", "internal")
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			fail("scoped DAO generation failed")
		}
		return
	}
	ctx := context.Background()
	config, err := g.Cfg().Get(ctx, "database.default.link")
	if err != nil {
		fail("database configuration unavailable")
	}
	db, err := sql.Open("mysql", strings.TrimPrefix(config.String(), "mysql:"))
	if err != nil {
		fail("database connection configuration invalid")
	}
	defer db.Close()
	conn, err := db.Conn(ctx)
	if err != nil {
		fail("database connection unavailable")
	}
	defer conn.Close()
	var lock int
	check(conn.QueryRowContext(ctx, "SELECT GET_LOCK('hotgo_shop_product_schema',10)").Scan(&lock))
	if lock != 1 {
		fail("another product schema migration holds the lock")
	}
	defer conn.ExecContext(ctx, "SELECT RELEASE_LOCK('hotgo_shop_product_schema')")
	oldCount, newCount := tablePresence(ctx, conn)
	fmt.Printf("table_state source=%d target=%d\n", oldCount, newCount)
	if oldCount == 1 && newCount == 1 {
		fail("target exists alongside source; refusing to overwrite or merge")
	}
	if oldCount+newCount != 1 {
		fail("exactly one source or migrated target BASE TABLE must exist")
	}
	table := oldProductTable
	if newCount == 1 {
		table = newProductTable
	}
	queries := []struct{ label, query string }{
		{"foreign_keys", "SELECT TABLE_SCHEMA,TABLE_NAME,COLUMN_NAME,REFERENCED_TABLE_NAME FROM information_schema.KEY_COLUMN_USAGE WHERE ((TABLE_SCHEMA=DATABASE() AND TABLE_NAME IN ('hg_web_products','hg_shop_product')) OR (REFERENCED_TABLE_SCHEMA=DATABASE() AND REFERENCED_TABLE_NAME IN ('hg_web_products','hg_shop_product'))) AND REFERENCED_TABLE_NAME IS NOT NULL"},
		{"views", "SELECT TABLE_SCHEMA,TABLE_NAME FROM information_schema.VIEWS WHERE VIEW_DEFINITION LIKE '%hg_web_products%' OR VIEW_DEFINITION LIKE '%hg_shop_product%'"},
		{"triggers", "SELECT TRIGGER_SCHEMA,TRIGGER_NAME,EVENT_OBJECT_TABLE FROM information_schema.TRIGGERS WHERE (TRIGGER_SCHEMA=DATABASE() AND EVENT_OBJECT_TABLE IN ('hg_web_products','hg_shop_product')) OR ACTION_STATEMENT LIKE '%hg_web_products%' OR ACTION_STATEMENT LIKE '%hg_shop_product%'"},
		{"routines", "SELECT ROUTINE_SCHEMA,ROUTINE_NAME FROM information_schema.ROUTINES WHERE ROUTINE_DEFINITION LIKE '%hg_web_products%' OR ROUTINE_DEFINITION LIKE '%hg_shop_product%'"},
		{"config_reference_keys", "SELECT id,`key`,LENGTH(value) AS value_length FROM hg_sys_config WHERE value LIKE '%hg_web_products%' OR value LIKE '%web_products%' OR value LIKE '%WebProducts%'"},
		{"addon_config_reference_keys", "SELECT id,`key`,LENGTH(value) AS value_length FROM hg_sys_addons_config WHERE value LIKE '%hg_web_products%' OR value LIKE '%web_products%' OR value LIKE '%WebProducts%'"},
		{"channel_models", "SELECT id,model FROM hg_web_channel WHERE model IN ('products','hg_web_products','web_products','WebProducts','hg_shop_product','ShopProduct')"},
		{"generator_values", "SELECT id,table_name,dao_name,var_name FROM hg_sys_gen_codes WHERE table_name IN ('hg_web_products','hg_shop_product') OR dao_name IN ('WebProducts','ShopProduct')"},
		{"generator_embedded_old_refs", "SELECT id FROM hg_sys_gen_codes WHERE CAST(options AS CHAR) LIKE '%hg_web_products%' OR CAST(options AS CHAR) LIKE '%web_products%' OR CAST(options AS CHAR) LIKE '%WebProducts%' OR CAST(master_columns AS CHAR) LIKE '%hg_web_products%' OR CAST(master_columns AS CHAR) LIKE '%web_products%' OR CAST(master_columns AS CHAR) LIKE '%WebProducts%'"},
	}
	for _, q := range queries {
		count, err := printQuery(ctx, conn, q.label, q.query)
		check(err)
		if *apply && count > 0 && (q.label == "views" || q.label == "triggers" || q.label == "routines" || q.label == "config_reference_keys" || q.label == "addon_config_reference_keys" || q.label == "generator_embedded_old_refs") {
			fail("unreviewed dependency or embedded metadata found; inspect before applying")
		}
	}
	before := snapshot(ctx, conn, table)
	fmt.Printf("before table=%s %+v\n", table, before)
	if !*apply {
		fmt.Println("inspect only; no schema or metadata changes")
		return
	}
	if oldCount == 1 {
		// MySQL RENAME TABLE has an implicit commit and cannot be rolled back.
		// Metadata updates are a separate resumable transaction after the rename.
		_, err = conn.ExecContext(ctx, "RENAME TABLE `hg_web_products` TO `hg_shop_product`")
		check(err)
		fmt.Println("atomic RENAME TABLE completed; no business rows inserted, deleted or updated")
	} else {
		fmt.Println("rename already applied; checking/resuming exact metadata updates")
	}
	tx, err := conn.BeginTx(ctx, nil)
	check(err)
	defer tx.Rollback()
	updates := []struct{ label, query string }{
		{"generator_table_name", "UPDATE hg_sys_gen_codes SET table_name='hg_shop_product' WHERE table_name='hg_web_products'"},
		{"generator_dao_name", "UPDATE hg_sys_gen_codes SET dao_name='ShopProduct' WHERE table_name='hg_shop_product' AND dao_name='WebProducts'"},
		{"channel_table_name", "UPDATE hg_web_channel SET model='hg_shop_product' WHERE model='hg_web_products'"},
	}
	for _, u := range updates {
		result, err := tx.ExecContext(ctx, u.query)
		check(err)
		n, err := result.RowsAffected()
		check(err)
		fmt.Printf("metadata_update %s=%d\n", u.label, n)
	}
	check(tx.Commit())
	after := snapshot(ctx, conn, newProductTable)
	fmt.Printf("after table=%s %+v\n", newProductTable, after)
	if !reflect.DeepEqual(before, after) {
		fail("post-rename full ID/data/schema or product-reference digest mismatch; investigate before business writes")
	}
	oldCount, newCount = tablePresence(ctx, conn)
	if oldCount != 0 || newCount != 1 {
		fail("post-rename table state mismatch")
	}
	fmt.Println("verified source absent, target present; all ordered IDs, all row values, schema and product references unchanged")
}

type productSnapshot struct {
	Count                                                            int64
	MinID, MaxID, IDSHA256, RowSHA256, SchemaSHA256, ReferenceSHA256 string
}

func snapshot(ctx context.Context, conn *sql.Conn, table string) productSnapshot {
	if table != oldProductTable && table != newProductTable {
		fail("invalid product table")
	}
	var result productSnapshot
	check(conn.QueryRowContext(ctx, "SELECT COUNT(*),COALESCE(MIN(id),0),COALESCE(MAX(id),0) FROM `"+table+"`").Scan(&result.Count, &result.MinID, &result.MaxID))
	result.IDSHA256 = digest(ctx, conn, "SELECT id FROM `"+table+"` ORDER BY id")
	result.RowSHA256 = digest(ctx, conn, "SELECT * FROM `"+table+"` ORDER BY id")
	// Exclude TABLE_NAME itself while checking structure and AUTO_INCREMENT.
	result.SchemaSHA256 = digest(ctx, conn, "SELECT COLUMN_NAME,COLUMN_TYPE,IS_NULLABLE,COLUMN_DEFAULT,EXTRA,COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=? ORDER BY ORDINAL_POSITION", table) + ":" + digest(ctx, conn, "SELECT INDEX_NAME,NON_UNIQUE,SEQ_IN_INDEX,COLUMN_NAME FROM information_schema.STATISTICS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=? ORDER BY INDEX_NAME,SEQ_IN_INDEX", table) + ":" + digest(ctx, conn, "SELECT ENGINE,AUTO_INCREMENT,TABLE_COLLATION FROM information_schema.TABLES WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=?", table)
	result.ReferenceSHA256 = digest(ctx, conn, "SELECT id,product_id FROM hg_shop_cart ORDER BY id") + ":" + digest(ctx, conn, "SELECT id,product_id FROM hg_shop_order_item ORDER BY id")
	return result
}

func digest(ctx context.Context, conn *sql.Conn, query string, args ...any) string {
	rows, err := conn.QueryContext(ctx, query, args...)
	check(err)
	defer rows.Close()
	columns, err := rows.Columns()
	check(err)
	h := sha256.New()
	for rows.Next() {
		values := make([]sql.RawBytes, len(columns))
		pointers := make([]any, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		check(rows.Scan(pointers...))
		for _, v := range values {
			if v == nil {
				_, _ = h.Write([]byte{0})
				continue
			}
			_, _ = h.Write([]byte{1})
			var size [8]byte
			binary.BigEndian.PutUint64(size[:], uint64(len(v)))
			_, _ = h.Write(size[:])
			_, _ = h.Write(v)
		}
	}
	check(rows.Err())
	return hex.EncodeToString(h.Sum(nil))
}

func tablePresence(ctx context.Context, conn *sql.Conn) (int, int) {
	var oldCount, newCount int
	check(conn.QueryRowContext(ctx, "SELECT COALESCE(SUM(TABLE_NAME=?),0),COALESCE(SUM(TABLE_NAME=?),0) FROM information_schema.TABLES WHERE TABLE_SCHEMA=DATABASE() AND TABLE_TYPE='BASE TABLE' AND TABLE_NAME IN (?,?)", oldProductTable, newProductTable, oldProductTable, newProductTable).Scan(&oldCount, &newCount))
	return oldCount, newCount
}

func printQuery(ctx context.Context, conn *sql.Conn, label, query string, args ...any) (int, error) {
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return 0, err
	}
	count := 0
	for rows.Next() {
		values := make([]sql.NullString, len(cols))
		pointers := make([]any, len(cols))
		for i := range values {
			pointers[i] = &values[i]
		}
		if err := rows.Scan(pointers...); err != nil {
			return count, err
		}
		fmt.Print(label)
		for i, c := range cols {
			fmt.Printf(" %s=%s", c, values[i].String)
		}
		fmt.Println()
		count++
	}
	fmt.Printf("%s rows=%d\n", label, count)
	return count, rows.Err()
}

func check(err error) {
	if err != nil {
		fail(err.Error())
	}
}
func fail(message string) { fmt.Fprintln(os.Stderr, message); os.Exit(1) }
