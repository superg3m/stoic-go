package ORM

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
	"strings"
)

// mysql
// sqlserver
// postgres
// sql_lite

func DeleteRecord(db *sqlx.DB, tableName string, model interface{}) error {
	fieldNames := Utility.GetStructMemberNames(model)
	if len(fieldNames) == 0 {
		return fmt.Errorf("no fields to use for DELETE condition in table '%s'", tableName)
	}

	var conditions []string
	values := Utility.GetStructValues(model)

	for _, fieldName := range fieldNames {
		conditions = append(conditions, fmt.Sprintf("%s = ?", fieldName))
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s",
		tableName,
		strings.Join(conditions, " AND "),
	)

	_, execErr := db.Exec(query, values...)
	if execErr != nil {
		return fmt.Errorf("failed to execute query: %w", execErr)
	}

	return nil
}

func UpdateRecord(db *sqlx.DB, tableName string, model interface{}) error {
	fieldNames := Utility.GetStructMemberNames(model)
	if len(fieldNames) <= 1 {
		return fmt.Errorf("not enough fields to construct an UPDATE statement for table '%s'", tableName)
	}

	values := Utility.GetStructValues(model)

	keyField := fieldNames[0] // Get the primary key
	updateFields := fieldNames[1:]
	updateValues := values[1:]
	keyValue := values[0]

	var updates []string
	for _, fieldName := range updateFields {
		updates = append(updates, fmt.Sprintf("%s = ?", fieldName))
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s = ?",
		tableName,
		strings.Join(updates, ", "),
		keyField,
	)

	_, execErr := db.Exec(query, append(updateValues, keyValue)...)
	if execErr != nil {
		return fmt.Errorf("failed to execute query: %w", execErr)
	}

	return nil
}

func getDBColumnNames(tableName string, fieldNames []string) []string {
	var ret []string
	for _, fieldName := range fieldNames {
		attr, exists := getAttribute(tableName, fieldName)
		Utility.Assert(exists)
		ret = append(ret, attr.ColumnName)
	}

	return ret
}

func InsertRecord(db *sqlx.DB, tableName string, model any) error {
	fieldNames := Utility.GetStructMemberNames(model)
	if len(fieldNames) == 0 {
		return fmt.Errorf("no fields to insert for table '%s'", tableName)
	}

	dbNames := getDBColumnNames(tableName, fieldNames)

	placeholders := make([]string, len(dbNames))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	values := Utility.GetStructValues(model)

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(dbNames, ", "),
		strings.Join(placeholders, ", "),
	)

	_, execErr := db.Exec(query, values...)
	if execErr != nil {
		return fmt.Errorf("failed to execute query: %w", execErr)
	}

	return nil
}

func GetDSN(dbEngine, host string, port int, user, password, dbname string) string {
	switch dbEngine {
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			user, password, host, port, dbname,
		)
	case "sqlite3":
		return dbname // For SQLite, `dbname` is the file path
	case "sqlserver":
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s:%d?database=%s",
			user, password, host, port, dbname,
		)
	default:
		return ""
	}
}

var db *sqlx.DB

func GetInstance() *sqlx.DB {
	Utility.Assert(db != nil)
	return db
}

func Connect(dbEngine, dsn string) *sqlx.DB {
	if db != nil {
		return db
	}

	dbNew, err := sqlx.Connect(dbEngine, dsn)
	Utility.AssertOnError(err)

	db = dbNew

	return db
}

func Close() {
	Utility.AssertMsg(db != nil, "Database Must have a active connection first before attempting to close")

	err := db.Close()
	if err != nil {
		return
	}

	db = nil
}
