package main

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
)

type SQLColumn struct {
	Name   string
	Type   string
	IsNull string
	IsKey  string
	Extra  string
}

type TableColumn struct {
	Name     string
	Type     string
	Flags    int
	StrFlags string
}

func (column *TableColumn) hasFlag(checkFlag int) bool {
	return column.Flags&checkFlag == 1
}

type Table struct {
	TableName    string
	TableColumns []TableColumn
	PrimaryKeys  []TableColumn
	UniqueKeys   []TableColumn

	RequireTimeInclude bool
}

func generateTable(tableName string, db *sqlx.DB, databaseName string) *Table {
	ret := &Table{}
	tableColumns, requireTimeInclude, err := FetchTableColumns(db, tableName, databaseName)
	Utility.AssertOnError(err)

	ret.RequireTimeInclude = requireTimeInclude
	ret.TableColumns = tableColumns
	ret.TableName = tableName
	for _, column := range ret.TableColumns {
		if column.hasFlag(IS_KEY) {
			ret.PrimaryKeys = append(ret.PrimaryKeys, column)
		} else if column.hasFlag(IS_UNIQUE) {
			ret.UniqueKeys = append(ret.UniqueKeys, column)
		}
	}

	return ret
}

func mapSQLTypeToGoType(sqlType string, isNull string) (string, bool) {
	if strings.Contains(sqlType, "varchar") {
		return "string", false
	}

	switch strings.ToLower(sqlType) {
	case "integer", "int":
		return "int", false
	case "varchar", "nvarchar":
		return "string", false
	case "timestamp", "datetime":
		if isNull == "NO" {
			return "time.Time", true
		}
		return "*time.Time", true
	case "tinyint":
		return "bool", false
	default:
		return "any", false
	}
}

func SQLColumnToTableColumn(sqlColumn SQLColumn) (TableColumn, bool) {
	requireTimeInclude := false

	tableColumn := TableColumn{}
	tableColumn.Name = sqlColumn.Name
	tableColumn.Type, requireTimeInclude = mapSQLTypeToGoType(sqlColumn.Type, sqlColumn.IsNull)
	tableColumn.Flags = generateFlags(sqlColumn.IsNull, sqlColumn.IsKey, sqlColumn.Extra)
	tableColumn.StrFlags = generateStrFlags(sqlColumn.IsNull, sqlColumn.IsKey, sqlColumn.Extra)

	return tableColumn, requireTimeInclude
}

func FetchTableColumns(db *sqlx.DB, tableName string, databaseName string) ([]TableColumn, bool, error) {
	var results []TableColumn
	sql := `
	SELECT
	COLUMN_NAME AS Field,
	COLUMN_TYPE AS Type,
	IS_NULLABLE AS IsNull,
	COLUMN_KEY AS IsKey,
	EXTRA
	FROM
	INFORMATION_SCHEMA.COLUMNS
	WHERE
	TABLE_NAME = ? AND TABLE_SCHEMA = ?
	ORDER BY 
    ORDINAL_POSITION
	`

	rows, err := db.Queryx(sql, tableName, databaseName)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	requireTimeInclude := false
	for rows.Next() {
		var sqlColumn SQLColumn
		pointers := Utility.GetStructMemberPointer(&sqlColumn)
		err := rows.Scan(pointers...)
		Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row into struct: %s", err))

		tableColumn, shouldRequireTimeInclude := SQLColumnToTableColumn(sqlColumn)
		if requireTimeInclude == false {
			requireTimeInclude = shouldRequireTimeInclude
		}

		results = append(results, tableColumn)
	}

	if err := rows.Err(); err != nil {
		return nil, false, fmt.Errorf("rows iteration error: %v", err)
	}

	return results, requireTimeInclude, nil
}
