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

// FetchTableRowALL retrieves all the rows for a given table.

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
