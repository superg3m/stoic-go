package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
	"strings"
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
	StrFlags []string
}

func (column *TableColumn) hasFlag(checkFlag int) bool {
	return column.Flags&checkFlag == 1
}

type Table struct {
	TableName    string
	TableColumns []TableColumn
	PrimaryKeys  []TableColumn
	UniqueKeys   []TableColumn
}

func generateTable(tableName string, db *sqlx.DB) *Table {
	ret := &Table{}
	tableColumns, err := FetchTableColumns(db, tableName)
	Utility.AssertOnError(err)

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

func mapSQLTypeToGoType(sqlType string, isNull string) string {
	if strings.Contains(sqlType, "varchar") {
		return "string"
	}

	switch strings.ToLower(sqlType) {
	case "integer", "int":
		return "int"
	case "varchar", "nvarchar":
		return "string"
	case "timestamp", "datetime":
		if isNull == "NO" {
			return "time.Time"
		}
		return "*time.Time"
	case "tinyint":
		return "bool"
	default:
		return "any"
	}
}

func SQLColumnToTableColumn(sqlColumn SQLColumn) TableColumn {
	tableColumn := TableColumn{}
	tableColumn.Name = sqlColumn.Name
	tableColumn.Type = mapSQLTypeToGoType(sqlColumn.Type, sqlColumn.IsNull)
	tableColumn.Flags = generateFlags(sqlColumn.IsNull, sqlColumn.IsKey, sqlColumn.Extra)
	tableColumn.StrFlags = generateStrFlags(sqlColumn.IsNull, sqlColumn.IsKey, sqlColumn.Extra)

	return tableColumn
}

func FetchTableColumns(db *sqlx.DB, tableName string) ([]TableColumn, error) {
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
	TABLE_NAME = ? AND TABLE_SCHEMA = 'stoic'
	ORDER BY 
    ORDINAL_POSITION
	`

	rows, err := db.Queryx(sql, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sqlColumn SQLColumn
		pointers := Utility.GetStructMemberPointer(&sqlColumn)
		err := rows.Scan(pointers...)
		Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row into struct: %s", err))

		tableColumn := SQLColumnToTableColumn(sqlColumn)

		results = append(results, tableColumn)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return results, nil
}
