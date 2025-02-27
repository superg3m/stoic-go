package main

/*
import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
	"strings"
)

func stringToBool(str string) bool {
	if strings.ToUpper(str) == "NO" {
		return false
	} else if strings.ToUpper(str) == "YES" {
		return true
	}

	return false
}

func extraToFlags(extra, isNull, isKey string) string {
	extraAttrs := strings.Fields(extra)
	var flags []string

	for _, attr := range extraAttrs {
		switch attr {
		case "auto_increment":
			flags = append(flags, "ORM.AUTO_INCREMENT")
		}
	}

	if isKey == "PRI" {
		flags = append(flags, "ORM.KEY")
	} else if isKey == "UNI" {
		flags = append(flags, "ORM.UNIQUE")
	}

	if stringToBool(isNull) {
		flags = append(flags, "ORM.NULLABLE")
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

func (r *TableRow) generatePrimaryKey() string {
	if r.IsKey == "PRI" {
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
				Type: mapSQLTypeToGoType(row.Type, row.IsNull),
			})
		}
	}

	return ret
}

func (r *TableRow) generateUnique() string {
	if r.IsKey == "UNI" {
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
				Type: mapSQLTypeToGoType(row.Type, row.IsNull),
			})
		}
	}
	return ret
}

// FetchTableRowALL retrieves all the rows for a given table.

func mapSQLTypeToGoType(sqlType string, isNull string) string {
	isNullable := stringToBool(isNull)

	if strings.Contains(sqlType, "varchar") {
		return "string"
	}

	switch strings.ToLower(sqlType) {
	case "integer", "int":
		return "int"
	case "varchar", "nvarchar":
		return "string"
	case "timestamp", "datetime":
		if !isNullable {
			return "time.Time"
		}
		return "*time.Time"
	case "tinyint":
		return "bool"
	default:
		return "any"
	}
}

func generateFromPrimaryKey(primaryKeys []PairData) string {
	var parts []string
	for i, pk := range primaryKeys {
		if i > 0 {
			parts = append(parts, "_")
		}
		parts = append(parts, pk.Name)
	}
	return strings.Join(parts, "")
}

func generatePrimaryKeyArgs(primaryKeys []PairData) string {
	var ret string
	for i, pk := range primaryKeys {
		if i > 0 {
			ret += ", "
		}

		ret += pk.Name
	}

	return ret
}

*/
