package main

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
	"html/template"
)

// ./cmd/bin/builder dsn password username dbname Table to build
// ./cmd/bin/builder "<TableName>"

// Parse through the database and look for a table called <TableName>.
// If files already exist, then just log: "File already exists for <TableName> table."

type TemplateDataType struct {
	TableName string
	//Columns                    []Attribute
	//PrimaryKeys                []PairData
	PrimaryKeyArgsWithTypes    string
	PrimaryKeyArgsWithoutTypes string
	//UniqueKeys                 []PairData
	UniqueKeyArgsWithTypes    string
	UniqueKeyArgsWithoutTypes string
	SafeHTML                  template.HTML
	FromPrimaryKey            string
	FromUniques               []string
	PrimaryKeyParams          string
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

	return

	/*
			'ClassName' => $table->name,
			'Columns' => $table->columns,
			'ColumnNameWithoutTypes => implode(", ", $ColumnArgsStrings),
			'ColumnNameWithTypes => implode(", ", $ColumnArgsStrings),

			'PrimaryKeys' => $table->primaryKeys,
			'PrimaryKeyArgsStrings' => implode(", ", $PrimaryKeyArgsStrings),
			'FromPrimaryKey' => implode("_", $PrimaryKeyArgsWithoutTypes),
		  	'PrimaryKeyArgs' => implode(", ", $PrimaryKeyArgsWithoutTypes),
		  	'PrimaryKeyArgsWithTypes' => implode(", ", $PrimaryKeyArgsWithTypes),

			'UniqueKeys' => $table->uniqueKeys,
			'FromUniqueKey' => implode("_", $UniqueKeyArgsWithoutTypes),
			'UniqueKeyArgsWithoutTypes' => implode(", ", $UniqueKeyArgsWithoutTypes),
			'UniqueKeyArgsWithTypes' => implode(", ", $UniqueKeyArgsWithTypes),
			'UniqueKeyArgs' => implode(", ", $UniqueKeyArgsWithoutTypes),
	*/

	/*
		templateData := TemplateDataType{
			TableName:     tableName,
			columns:       columns,
			ColumnArgsStrings:

			Attributes:    attributes,
			PrimaryKeys:   primaryKeys,
			PrimaryKeyArg: primaryKeyArg,
			UniqueKeys:    uniqueKeys,
			SafeHTML:      template.HTML(`<`),
			// Precompute FromPrimaryKey
			FromPrimaryKey:   generateFromPrimaryKey(primaryKeys),
			PrimaryKeyParams: strings.Join(primaryKeys),
		}

		tmplFile := "./cmd/bin/templates/cls.tmpl"
		tmpl, err := template.ParseFiles(tmplFile)
		Utility.AssertOnError(err)

		dirName := fmt.Sprintf("./inc/%s", tableName)

		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0755)
			Utility.AssertOnError(err)
		}

		filePtr, err := os.Create(fmt.Sprintf("%s/%s.cls.go", dirName, tableName))
		Utility.AssertOnError(err)

		err = tmpl.Execute(filePtr, templateData)
		Utility.AssertOnError(err)

		// --------------------------------------------------------

		tmplFile = "./cmd/bin/templates/api.tmpl"
		tmpl, err = template.ParseFiles(tmplFile)
		Utility.AssertOnError(err)

		filePtr, err = os.Create(fmt.Sprintf("./API/0.1/%s.api.go", tableName))
		Utility.AssertOnError(err)

		err = tmpl.Execute(filePtr, templateData)
		Utility.AssertOnError(err)

		// --------------------------------------------------------

		tmplFile = "./cmd/bin/templates/crud.tmpl"
		tmpl, err = template.ParseFiles(tmplFile)
		Utility.AssertOnError(err)

		filePtr, err = os.Create(fmt.Sprintf("%s/%s.crud.go", dirName, tableName))
		Utility.AssertOnError(err)

		err = tmpl.Execute(filePtr, templateData)
		Utility.AssertOnError(err)

		// --------------------------------------------------------

		tmplFile = "./cmd/bin/templates/model.tmpl"
		tmpl, err = template.ParseFiles(tmplFile)
		Utility.AssertOnError(err)

		filePtr, err = os.Create(fmt.Sprintf("%s/%s.model.go", dirName, tableName))
		Utility.AssertOnError(err)

		err = tmpl.Execute(filePtr, templateData)
		Utility.AssertOnError(err)
	
	*/
}
