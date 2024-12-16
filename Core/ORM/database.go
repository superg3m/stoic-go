package ORM

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
)

// mysql
// sqlserver
// postgres
// sql_lite

func DeleteRecord[T InterfaceCRUD](db *sqlx.DB, tableName string, model *T) (sql.Result, error) {
	fieldNames := getDBColumnNames(tableName, *model)
	Utility.Assert(len(fieldNames) > 0)

	var conditions []string
	values := Utility.GetStructValues(*model)

	for _, fieldName := range fieldNames {
		conditions = append(conditions, fmt.Sprintf("%s = ?", fieldName))
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s",
		tableName,
		strings.Join(conditions, " AND "),
	)

	result, execErr := db.Exec(query, values...)
	if execErr != nil {
		return result, fmt.Errorf("failed to execute query: %w", execErr)
	}

	return result, nil
}

func UpdateRecord[T InterfaceCRUD](db *sqlx.DB, tableName string, model *T) (sql.Result, error) {
	fieldNames := getDBColumnNames(tableName, *model)
	Utility.Assert(len(fieldNames) > 0)

	values := Utility.GetStructValues(*model)

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

	result, execErr := db.Exec(query, append(updateValues, keyValue)...)
	if execErr != nil {
		return result, fmt.Errorf("failed to execute query: %w", execErr)
	}

	return result, nil
}

func updateStoicModel[T InterfaceCRUD](model *T) {
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("model must be a pointer to a struct")
	}

	structValue := v.Elem()

	stoicModelField := structValue.FieldByName("StoicModel")
	if !stoicModelField.IsValid() {
		Utility.AssertMsg(false, "Embedded StoicModel is missing from the struct")
	}

	if stoicModelField.Kind() == reflect.Struct && stoicModelField.CanSet() {
		isCreatedField := stoicModelField.FieldByName("IsCreated")
		if isCreatedField.IsValid() && isCreatedField.CanSet() && isCreatedField.Kind() == reflect.Bool {
			isCreatedField.SetBool(true)
		} else {
			Utility.AssertMsg(false, "IsCreated field in StoicModel is missing, not settable, or not of type bool")
		}
	} else {
		Utility.AssertMsg(false, "StoicModel is not a struct or is not settable")
	}
}

func updateIDField[T InterfaceCRUD](model *T, id int64) {
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("model must be a pointer to a struct")
	}

	structValue := v.Elem()
	field := structValue.FieldByName("ID")
	if field.IsValid() && field.CanSet() && field.Kind() == reflect.Int {
		field.SetInt(id)
	} else {
		Utility.AssertMsg(false, "ID field is missing, not settable, or not of type int")
	}
}

func InsertRecord[T InterfaceCRUD](db *sqlx.DB, tableName string, model *T) (sql.Result, error) {
	fieldNames := Utility.GetStructMemberNames(*model)
	Utility.Assert(len(fieldNames) > 0)

	dbNames := getDBColumnNames(tableName, *model)
	placeholders := make([]string, len(dbNames))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	values := Utility.GetStructValues(*model)

	var newDbNames []string
	var newPlaceholders []string
	var newValues []interface{}

	for i, fieldName := range fieldNames {
		attribute := globalTable[tableName][fieldName]
		if attribute.isAutoIncrement() {
			continue
		}
		newDbNames = append(newDbNames, dbNames[i])
		newPlaceholders = append(newPlaceholders, placeholders[i])
		newValues = append(newValues, values[i])
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(newDbNames, ", "),
		strings.Join(newPlaceholders, ", "),
	)

	result, execErr := db.Exec(query, newValues...)
	if execErr != nil {
		return result, fmt.Errorf("failed to execute query: %s\nError: %w", query, execErr)
	}

	return result, nil
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

func getDBColumnNames[T InterfaceCRUD](tableName string, model T) []string {
	var ret []string

	fieldNames := Utility.GetStructMemberNames(model)
	for _, fieldName := range fieldNames {
		attr, exists := getAttribute(tableName, fieldName)
		Utility.Assert(exists)
		ret = append(ret, attr.ColumnName)
	}

	return ret
}
