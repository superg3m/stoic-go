package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
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
	TableName   string
	Attributes  []Attribute
	PrimaryKeys []PairData
	UniqueKeys  []PairData
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

	tableName := "User"
	db := ORM.GetInstance()

	attributes, err := getAttributes(db, tableName)
	Utility.AssertOnError(err)

	primaryKeys, err := getPrimaryKeys(db, tableName)
	Utility.AssertOnError(err)

	uniqueKeys, err := getUniqueKeys(db, tableName)
	Utility.AssertOnError(err)

	templateData := TemplateDataType{
		TableName:   tableName,
		Attributes:  attributes,
		PrimaryKeys: primaryKeys,
		UniqueKeys:  uniqueKeys,
	}
	Utility.AssertOnError(err)

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

func getAttributes(db *sqlx.DB, tableName string) ([]Attribute, error) {
	attributes := []Attribute{}
	query := `SHOW COLUMNS FROM ` + tableName

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var columnName, columnType, isNullable, key, defaultValue, extra sql.NullString
		if err := rows.Scan(&columnName, &columnType, &isNullable, &key, &defaultValue, &extra); err != nil {
			return nil, err
		}

		flags := []string{}
		if isNullable.Valid && isNullable.String == "NO" {
			flags = append(flags, "NOT NULL")
		}
		if key.Valid && key.String == "PRI" {
			flags = append(flags, "PRIMARY KEY")
		}
		if key.Valid && key.String == "UNI" {
			flags = append(flags, "UNIQUE")
		}

		attributes = append(attributes, Attribute{
			Name:   strings.Title(columnName.String),
			Type:   mapSQLTypeToGoType(columnType.String, flags),
			Column: columnName.String,
			Flags:  strings.Join(flags, "|"),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attributes, nil
}

func getPrimaryKeys(db *sqlx.DB, tableName string) ([]PairData, error) {
	primaryKeys := []PairData{}
	query := `SHOW KEYS FROM ` + tableName + ` WHERE Key_name = 'PRIMARY'`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table, columnName, indexName, collation, cardinality, subPart, packed, nullAllowed, indexType, comment, indexComment sql.NullString
		var nonUnique, seqInIndex sql.NullInt64
		// Update the Scan to match the expected 15 columns returned by SHOW KEYS
		if err := rows.Scan(
			&table, &nonUnique, &indexName, &seqInIndex, &columnName,
			&collation, &cardinality, &subPart, &packed, &nullAllowed,
			&indexType, &comment, &indexComment); err != nil {
			return nil, err
		}

		primaryKeys = append(primaryKeys, PairData{
			Name: columnName.String,
			Type: "PRIMARY KEY", // Replace with mapSQLTypeToGoType if needed
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return primaryKeys, nil
}

func getUniqueKeys(db *sqlx.DB, tableName string) ([]PairData, error) {
	uniqueKeys := []PairData{}
	query := `SHOW KEYS FROM ` + tableName + ` WHERE Non_unique = 0 AND Key_name != 'PRIMARY'`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table, columnName, indexName, collation, cardinality, subPart, packed, nullAllowed, indexType, comment, indexComment sql.NullString
		var nonUnique, seqInIndex sql.NullInt64
		if err := rows.Scan(&table, &nonUnique, &indexName, &seqInIndex, &columnName, &collation, &cardinality, &subPart, &packed, &nullAllowed, &indexType, &comment, &indexComment); err != nil {
			return nil, err
		}

		uniqueKeys = append(uniqueKeys, PairData{
			Name: columnName.String,
			Type: "UNIQUE", // Replace with mapSQLTypeToGoType if needed
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return uniqueKeys, nil
}

func mapSQLTypeToGoType(sqlType string, flags []string) string {
	isNotNull := containsFlag(flags, "NOT NULL")

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

func containsFlag(flags []string, target string) bool {
	for _, flag := range flags {
		if flag == target {
			return true
		}
	}
	return false
}
