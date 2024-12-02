package main

import (
	"github.com/superg3m/stoic-go/core/Utility"
	"html/template"
	"os"
)

// ./cmd/bin/builder dsn password username dbname Table to build
// ./cmd/bin/builder "<TableName>"

// Parse through the database and look for a table called <TableName>.
// If files already exist, then just log: "File already exists for <TableName> table."

type Attribute struct {
	Name     string
	Type     string
	DBColumn string
	Flags    string
}

type TemplateDataType struct {
	TableName  string
	Attributes []Attribute
}

func main() {
	tableName := "User"
	attributes := []Attribute{
		{"ID", "int", "user_id", "ORM.PRIMARY_KEY"},
		{"Email", "string", "email_address", "ORM.NULLABLE|ORM.UPDATABLE"},
		{"Joined", "time.Time", "joined_at", "ORM.NULLABLE"},
	}

	templateData := TemplateDataType{
		TableName:  tableName,
		Attributes: attributes,
	}

	tmplFile := "CLS.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	filePtr, err := os.Create("./user.cls.go")
	Utility.AssertOnError(err)

	err = tmpl.Execute(filePtr, templateData)
	if err != nil {
		panic(err)
	}
}
