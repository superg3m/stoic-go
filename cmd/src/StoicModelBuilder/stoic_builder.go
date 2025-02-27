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
	TableName           string
	Columns             []TableColumn
	ColumnNames         []string
	ColumnArgs          string
	ColumnArgsWithTypes string

	PrimaryKeys             []TableColumn
	PrimaryKeyNames         []string
	PrimaryKeyArgs          string
	PrimaryKeyArgsWithTypes string
	FromPrimaryKey          string

	UniqueKeys          []TableColumn
	UniqueNames         []string
	UniqueArgs          string
	UniqueArgsWithTypes string
	FromUniques         []string
	SafeHTML            template.HTML
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

	table := generateTable(tableName, db)
	Utility.Assert(table != nil)

	var columnNames []string
	var columnNamesWithTypes []string
	var primaryKeyNames []string
	var primaryKeyNamesWithTypes []string
	var uniqueNames []string
	var uniqueNamesWithTypes []string

	for _, column := range table.TableColumns {
		columnNames = append(columnNames, column.Name)
		columnNamesWithTypes = append(columnNamesWithTypes, fmt.Sprintf("%s %s", column.Name, column.Type))

		if column.hasFlag(IS_KEY) {
			primaryKeyNames = append(primaryKeyNames, column.Name)
			primaryKeyNamesWithTypes = append(primaryKeyNamesWithTypes, fmt.Sprintf("%s %s", column.Name, column.Type))
		} else if column.hasFlag(IS_UNIQUE) {
			uniqueNames = append(uniqueNames, column.Name)
			uniqueNamesWithTypes = append(uniqueNamesWithTypes, fmt.Sprintf("%s %s", column.Name, column.Type))
		}
	}

	columnArgs := strings.Join(columnNames, ", ")
	columnArgsWithTypes := strings.Join(columnNamesWithTypes, ", ")

	primaryKeyArgs := strings.Join(primaryKeyNames, ", ")
	primaryKeyArgsWithTypes := strings.Join(primaryKeyNamesWithTypes, ", ")
	fromPrimaryKeyMethodName := strings.Join(primaryKeyNames, "_")

	uniqueArgs := strings.Join(primaryKeyNames, ", ")
	uniqueArgsWithTypes := strings.Join(primaryKeyNamesWithTypes, ", ")

	templateData := TemplateDataType{
		TableName: tableName,

		Columns:             table.TableColumns,
		ColumnNames:         columnNames,
		ColumnArgs:          columnArgs,
		ColumnArgsWithTypes: columnArgsWithTypes,

		PrimaryKeys:             table.PrimaryKeys,
		PrimaryKeyArgs:          primaryKeyArgs,
		PrimaryKeyArgsWithTypes: primaryKeyArgsWithTypes,
		FromPrimaryKey:          fromPrimaryKeyMethodName,

		UniqueKeys:          table.UniqueKeys,
		UniqueNames:         uniqueNames,
		UniqueArgs:          uniqueArgs,
		UniqueArgsWithTypes: uniqueArgsWithTypes,

		SafeHTML: template.HTML(`<`),
	}

	tmplFile := "./cmd/bin/templates/cls.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	Utility.AssertOnError(err)
	dirName := fmt.Sprintf("./inc/%s", tableName)
	if _, err = os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0755)
		Utility.AssertOnError(err)
	}

	{
		clsFile := fmt.Sprintf("%s/%s.cls.go", dirName, tableName)
		if _, err = os.Stat(clsFile); os.IsNotExist(err) {
			filePtr, err := os.Create(clsFile)
			Utility.AssertOnError(err)

			err = tmpl.Execute(filePtr, templateData)
			Utility.AssertOnError(err)
			err = filePtr.Close()
			if err != nil {
				return
			}
		}
	}

	// --------------------------------------------------------

	{
		tmplFile = "./cmd/bin/templates/api.tmpl"
		tmpl, err = template.ParseFiles(tmplFile)
		Utility.AssertOnError(err)

		apiFile := fmt.Sprintf("./API/0.1/%s.api.go", tableName)
		if _, err = os.Stat(apiFile); os.IsNotExist(err) {
			filePtr, err := os.Create(apiFile)
			Utility.AssertOnError(err)

			err = tmpl.Execute(filePtr, templateData)
			Utility.AssertOnError(err)
			err = filePtr.Close()
			if err != nil {
				return
			}
		}
	}

	// --------------------------------------------------------

	{
		tmplFile = "./cmd/bin/templates/crud.tmpl"
		tmpl, err = template.ParseFiles(tmplFile)
		Utility.AssertOnError(err)

		crudFile := fmt.Sprintf("%s/%s.crud.go", dirName, tableName)
		if _, err = os.Stat(crudFile); os.IsNotExist(err) {
			filePtr, _ := os.Create(crudFile)

			err = tmpl.Execute(filePtr, templateData)
			Utility.AssertOnError(err)
			err = filePtr.Close()
			if err != nil {
				return
			}
		}
	}

	// --------------------------------------------------------

	{
		tmplFile = "./cmd/bin/templates/meta.tmpl"
		tmpl, err = template.ParseFiles(tmplFile)
		Utility.AssertOnError(err)

		filePtr, _ := os.Create(fmt.Sprintf("%s/%s.meta.go", dirName, tableName))

		err = tmpl.Execute(filePtr, templateData)
		Utility.AssertOnError(err)
		err = filePtr.Close()
		if err != nil {
			return
		}
	}

}
