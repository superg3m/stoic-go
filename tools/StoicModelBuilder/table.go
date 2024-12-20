package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
	"strings"
)

type TableRow struct {
	Field  string
	Type   string
	IsNull string
	IsKey  string
	Extra  string
}

type Table struct {
	TableName string
	TableRows []TableRow
}

type Attribute struct {
	Name   string
	Type   string
	Column string
	Flags  string
}

type PairData struct {
	Name string
	Type string
}

func stringToBool(str string) bool {
	if strings.ToUpper(str) == "NO" {
		return false
	} else if strings.ToUpper(str) == "YES" {
		return true
	}

	return false
}

func extraToFlags(extra string) string {
	extraAttrs := strings.Fields(extra)
	var flags []string

	for _, attr := range extraAttrs {
		switch attr {
		case "auto_increment":
			flags = append(flags, "ORM.AUTO_INCREMENT")
		case "unique":
			flags = append(flags, "ORM.UNIQUE")
		case "nullable":
			flags = append(flags, "ORM.NULLABLE")
		case "key":
			flags = append(flags, "ORM.KEY")
		}
	}

	return strings.Join(flags, " | ")
}

func extraHas(extra string, has string) bool {
	extraAttrs := strings.Fields(extra)

	for _, attr := range extraAttrs {
		if attr == has {
			return true
		}
	}

	return false
}

func (t *Table) generateTable(tableName string, db *sqlx.DB) error {
	tableRows, err := FetchTableRowALL(db, tableName)
	if err != nil {
		return err
	}
	t.TableRows = tableRows
	return nil
}

func (r *TableRow) generateAttribute() Attribute {
	return Attribute{
		Name:   r.Field,
		Type:   mapSQLTypeToGoType(r.Type, r.Extra),
		Column: r.Field,
		Flags:  extraToFlags(r.Extra),
	}
}

func (t *Table) generateAttributes() []Attribute {
	var ret []Attribute
	for _, row := range t.TableRows {
		ret = append(ret, row.generateAttribute())
	}
	return ret
}

func (r *TableRow) generatePrimaryKey() string {
	if stringToBool(r.IsKey) {
		return r.Field
	}
	return ""
}

func (t *Table) generatePrimaryKeys() []PairData {
	var ret []PairData
	for _, row := range t.TableRows {
		pKey := row.generatePrimaryKey()
		if pKey != "" {
			ret = append(ret, PairData{
				Name: row.Field,
				Type: row.Type,
			})
		}
	}
	return ret
}

func (r *TableRow) generateUnique() string {
	if extraHas(r.Extra, "unique") || extraHas(r.Extra, "UNIQUE") {
		return r.Field
	}
	return ""
}

func (t *Table) generateUniques() []PairData {
	var ret []PairData
	for _, row := range t.TableRows {
		unique := row.generateUnique()
		if unique != "" {
			ret = append(ret, PairData{
				Name: row.Field,
				Type: mapSQLTypeToGoType(row.Type, row.Type),
			})
		}
	}
	return ret
}

// FetchTableRowALL retrieves all the rows for a given table.
func FetchTableRowALL(db *sqlx.DB, tableName string) ([]TableRow, error) {
	var results []TableRow
	query := `
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

	rows, err := db.Queryx(query, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row TableRow
		pointers := Utility.GetStructMemberPointer(&row)
		err := rows.Scan(pointers...)
		Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row into struct: %s", err))

		results = append(results, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return results, nil
}

func mapSQLTypeToGoType(sqlType string, flags string) string {
	isNotNull := extraHas(flags, "NOT NULL")

	switch sqlType {
	case "integer", "int":
		return "int"
	case "varchar", "nvarchar":
		return "string"
	case "timestamp", "datetime":
		if isNotNull {
			return "time.Time"
		}
		return "*time.Time"
	case "tinyint":
		return "bool"
	default:
		return "any"
	}
}
