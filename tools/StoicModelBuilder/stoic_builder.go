package main

import (
	"github.com/superg3m/stoic-go/Core/Utility"
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

type PairData struct {
	Name string
	Type string
}

type TemplateDataType struct {
	TableName   string
	Attributes  []Attribute
	PrimaryKeys []PairData
	UniqueKeys  []PairData
}

func main() {
	/*
		siteSettings := Utility.GetSiteSettings()
		siteSettings = siteSettings["settings"].(map[string]any)
		DB_ENGINE := Utility.CastAny[string](siteSettings["dbEngine"])
		HOST := Utility.CastAny[string](siteSettings["dbHost"])
		PORT := Utility.CastAny[int](siteSettings["dbPort"])
		USER := Utility.CastAny[string](siteSettings["dbUser"])
		PASSWORD := Utility.CastAny[string](siteSettings["dbPass"])
		DBNAME := Utility.CastAny[string](siteSettings["dbName"])

		dsn := ORM.GetDSN(DB_ENGINE, HOST, PORT, USER, PASSWORD, DBNAME)
		ORM.Connect(DB_ENGINE, dsn)
		defer ORM.Close()


		// Describe the database now for a specifc table use command line
		// Or mabye use input()
	*/

	tableName := "User"
	attributes := []Attribute{
		{"ID", "int", "user_id", "ORM.KEY"},
		{"Username", "string", "userName", "ORM.KEY"},
		{"Joined", "time.Time", "joined", "ORM.UNIQUE"},
		{"Email", "string", "email_address", "ORM.UNIQUE|ORM.UPDATABLE"},
	}
	primaryKeys := []PairData{
		{"ID", "int"},
		{"Username", "string"},
	}
	uniqueKeys := []PairData{
		{"Email", "string"},
		{"Joined", "time.Time"},
	}

	templateData := TemplateDataType{
		TableName:   tableName,
		Attributes:  attributes,
		PrimaryKeys: primaryKeys,
		UniqueKeys:  uniqueKeys,
	}

	tmplFile := "cls.tmpl"
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

	// --------------------------------------------------------

	tmplFile = "api.tmpl"
	tmpl, err = template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	filePtr, err = os.Create("./user.api.go")
	Utility.AssertOnError(err)

	err = tmpl.Execute(filePtr, templateData)
	if err != nil {
		panic(err)
	}

	// --------------------------------------------------------

	tmplFile = "crud.tmpl"
	tmpl, err = template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	filePtr, err = os.Create("./user.crud.go")
	Utility.AssertOnError(err)

	err = tmpl.Execute(filePtr, templateData)
	if err != nil {
		panic(err)
	}
}
