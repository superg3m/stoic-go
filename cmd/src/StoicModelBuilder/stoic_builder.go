package main

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
	"html/template"
	"os"
	"strings"
)

// ./cmd/bin/builder dsn password username dbname Table to build
// ./cmd/bin/builder "<TableName>"

// Parse through the database and look for a table called <TableName>.
// If files already exist, then just log: "File already exists for <TableName> table."

type TemplateDataType struct {
	TableName        string
	Attributes       []Attribute
	PrimaryKeys      []PairData
	UniqueKeys       []PairData
	SafeHTML         template.HTML
	FromPrimaryKey   string
	FromUniques      []string
	PrimaryKeyParams string
}

func main() {
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

	tableName := ""
	fmt.Print("Enter TableName: ")
	_, err := fmt.Scanln(&tableName)
	fmt.Println("")
	Utility.AssertOnError(err)

	db := ORM.GetInstance()

	table := Table{
		TableName: tableName,
	}

	err = table.generateTable(tableName, db)
	Utility.AssertOnError(err)

	attributes := table.generateAttributes()
	primaryKeys := table.generatePrimaryKeys()
	uniqueKeys := table.generateUniques()

	templateData := TemplateDataType{
		TableName:   tableName,
		Attributes:  attributes,
		PrimaryKeys: primaryKeys,
		UniqueKeys:  uniqueKeys,
		SafeHTML:    template.HTML(`<`),
		// Precompute FromPrimaryKey
		FromPrimaryKey:   generateFromPrimaryKey(primaryKeys),
		PrimaryKeyParams: generatePrimaryKeyArgs(primaryKeys),
	}

	tmplFile := "./templates/cls.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	Utility.AssertOnError(err)

	dirName := fmt.Sprintf("../../inc/%s", tableName)

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0755)
		Utility.AssertOnError(err)
	}

	filePtr, err := os.Create(fmt.Sprintf("%s/%s.cls.go", dirName, strings.ToLower(tableName)))
	Utility.AssertOnError(err)

	err = tmpl.Execute(filePtr, templateData)
	Utility.AssertOnError(err)

	// --------------------------------------------------------

	tmplFile = "./templates/api.tmpl"
	tmpl, err = template.ParseFiles(tmplFile)
	Utility.AssertOnError(err)

	filePtr, err = os.Create(fmt.Sprintf("../../API/0.1/%s.api.go", strings.ToLower(tableName)))
	Utility.AssertOnError(err)

	err = tmpl.Execute(filePtr, templateData)
	Utility.AssertOnError(err)

	// --------------------------------------------------------

	tmplFile = "./templates/crud.tmpl"
	tmpl, err = template.ParseFiles(tmplFile)
	Utility.AssertOnError(err)

	filePtr, err = os.Create(fmt.Sprintf("%s/%s.crud.go", dirName, strings.ToLower(tableName)))
	Utility.AssertOnError(err)

	err = tmpl.Execute(filePtr, templateData)
	Utility.AssertOnError(err)
}
