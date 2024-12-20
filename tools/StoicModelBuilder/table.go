package main

import (
	"strings"
)

// | Field          | Type         | IS_NULLABLE | Key | Default | EXTRA          |
type TableRow struct {
	Field  string
	Type   string
	isNull bool
	isKey  string
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

func (r *TableRow) generateAttribute() Attribute {
	ret := Attribute{
		Name:   r.Field,
		Type:   r.Type,
		Column: r.Field,
		Flags:  extraToFlags(r.Extra),
	}

	return ret
}

func (t *Table) generateAttributes() []Attribute {
	ret := []Attribute{}
	for _, row := range t.TableRows {
		row.generateAttribute()
	}

	return ret
}

func (t *Table) generatePrimaryKeys() {

}
func (t *Table) generateUniques() {

}
