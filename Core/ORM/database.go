package ORM

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
)

// mysql
// sqlserver
// postgres
// sql_lite

func CreateRecord[T InterfaceCRUD](db *sqlx.DB, model *T) (sql.Result, error) {
	tableName := getModelTableName(*model)
	fieldNames := getModelMemberNames(*model)
	Utility.Assert(len(fieldNames) > 0)

	dbNames := getDBColumnNames(tableName, *model)
	placeholders := make([]string, len(dbNames))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	values, err := getCanonicalValues(*model, fieldNames)
	if err != nil {
		return nil, err
	}

	var newDbNames []string
	var newPlaceholders []string
	var newValues []any

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

func ReadRecord[T InterfaceCRUD](db *sqlx.DB, model *T) error {
	tableName := getModelTableName(*model)
	fieldNames := getModelMemberNames(*model)
	Utility.Assert(len(fieldNames) > 0)

	pKeyQuery, uniqueQueries, _ := buildSQLReadQueries(db, *model)

	// ----------------------------------------------------------

	{
		pPointer := getPrimaryKeyPointers(tableName, *model)
		temp, err := Fetch[T](pKeyQuery, pPointer...)
		Utility.AssertOnError(err)
		if err == nil {
			*model = temp
			return nil
		}
	}

	// ----------------------------------------------------------

	{
		uPointer := getUniquePointers(tableName, *model)
		for i, pointer := range uPointer {
			query := uniqueQueries[i]
			temp, err := Fetch[T](query, pointer)
			Utility.AssertOnError(err)
			if err == nil {
				*model = temp
				return nil
			}
		}
	}

	return errors.New("failed to fetch record")
}

func UpdateRecord[T InterfaceCRUD](db *sqlx.DB, model *T) (sql.Result, error) {
	tableName := getModelTableName(*model)
	fieldNames := getDBColumnNames(tableName, *model)
	Utility.Assert(len(fieldNames) > 0)

	values := getModelValues(*model)

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

func DeleteRecord[T InterfaceCRUD](db *sqlx.DB, model *T) (sql.Result, error) {
	tableName := getModelTableName(*model)
	fieldNames := getDBColumnNames(tableName, *model)
	Utility.Assert(len(fieldNames) > 0)

	var conditions []string
	values := getModelValues(*model)

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
		return result, fmt.Errorf("failed to execute query: %s", execErr)
	}

	rows, err2 := result.RowsAffected()
	if rows == 0 {
		return nil, fmt.Errorf("failed to execute query: Zero rows affected")
	}

	if err2 != nil {
		return nil, fmt.Errorf("failed to execute query: %s", err2)
	}

	err := Utility.SetToNil[T](model)
	if err != nil {
		return nil, err
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

	fieldNames := getModelMemberNames(model)
	for _, fieldName := range fieldNames {
		attr, exists := getAttribute(tableName, fieldName)
		Utility.Assert(exists)
		ret = append(ret, attr.ColumnName)
	}

	return ret
}

func getPrimaryKeyNames[T InterfaceCRUD](tableName string, model T) []string {
	var ret []string

	names := getModelMemberNames(model)
	attributes, _ := GetAttributes(tableName)
	for _, name := range names {
		attribute := attributes[name]
		if attribute.isPrimaryKey() {
			ret = append(ret, name)
		}
	}

	return ret
}

func getUniqueNames[T InterfaceCRUD](tableName string, model T) []string {
	var ret []string

	names := Utility.GetStructMemberNames(model, excludeList...)
	attributes, _ := GetAttributes(tableName)
	for _, name := range names {
		attribute := attributes[name]
		if attribute.isUnique() {
			ret = append(ret, name)
		}
	}

	return ret
}

func getPrimaryKeyDBNames[T InterfaceCRUD](tableName string, model T) []string {
	stackModel := Utility.DereferencePointer(model)
	var ret []string

	names := Utility.GetStructMemberNames(stackModel, excludeList...)
	attributes, _ := GetAttributes(tableName)
	for _, name := range names {
		attribute := attributes[name]
		if attribute.isPrimaryKey() {
			ret = append(ret, name)
		}
	}

	return ret
}

func getUniqueDBNames[T InterfaceCRUD](tableName string, model T) []string {
	var ret []string

	names := getModelMemberNames(model)
	attributes, _ := GetAttributes(tableName)
	for _, name := range names {
		attribute := attributes[name]
		if attribute.isUnique() {
			ret = append(ret, name)
		}
	}

	return ret
}

func getPrimaryKeyPointers[T InterfaceCRUD](tableName string, model T) []any {
	var ret []any

	pointers := Utility.GetStructMemberPointer(model, excludeList...)
	names := getModelMemberNames(model)
	attributes, _ := GetAttributes(tableName)
	for i, pointer := range pointers {
		name := names[i]
		attribute := attributes[name]
		if attribute.isPrimaryKey() {
			ret = append(ret, pointer)
		}
	}

	return ret
}

func getUniquePointers[T InterfaceCRUD](tableName string, model T) []any {
	var ret []any

	pointers := Utility.GetStructMemberPointer(model, excludeList...)
	names := getModelMemberNames(model)
	attributes, _ := GetAttributes(tableName)
	for i, pointer := range pointers {
		name := names[i]
		attribute := attributes[name]
		if attribute.isUnique() {
			ret = append(ret, pointer)
		}
	}

	return ret
}

func buildSQLReadQueries[T InterfaceCRUD](db *sqlx.DB, model T) (primaryQuery string, uniqueQueries []string, err error) {
	tableName := getModelTableName(model)
	fieldNames := getModelMemberNames(model)
	Utility.Assert(len(fieldNames) > 0)

	primaryKeyNames := getPrimaryKeyDBNames(tableName, model)
	if len(primaryKeyNames) == 0 {
		return "", nil, fmt.Errorf("no primary keys defined for table: %s", tableName)
	}

	primaryPlaceholders := make([]string, len(primaryKeyNames))
	for i := range primaryPlaceholders {
		primaryPlaceholders[i] = "?"
	}
	primaryQuery = fmt.Sprintf(
		"SELECT * FROM %s WHERE %s",
		tableName,
		strings.Join(primaryKeyNames, " = ? AND ")+" = ?",
	)

	uniqueKeyGroups := getUniqueDBNames(tableName, model)
	for _, uniqueKey := range uniqueKeyGroups {
		uniqueQuery := fmt.Sprintf(
			"SELECT * FROM %s WHERE %s = ?",
			tableName,
			uniqueKey,
		)
		uniqueQueries = append(uniqueQueries, uniqueQuery)
	}

	return primaryQuery, uniqueQueries, nil
}

func getCanonicalValues[T InterfaceCRUD](model T, fieldNames []string) ([]any, error) {
	stackModel := Utility.DereferencePointer(model)

	var formattedTimeValues []any
	types := Utility.GetStructMemberTypes(stackModel, excludeList...)

	for i, fieldName := range fieldNames {
		fieldType, exists := types[fieldName]
		Utility.Assert(exists)
		if fieldType == "time.Time" || fieldType == "*time.Time" {
			value := reflect.ValueOf(model).Elem().FieldByName(fieldName)
			if value.IsValid() && value.Kind() == reflect.Struct && value.Type() == reflect.TypeOf(time.Time{}) {
				formattedTime := value.Interface().(time.Time).Format(time.DateTime)
				formattedTimeValues = append(formattedTimeValues, formattedTime)
			}
		} else {
			originalValue := Utility.GetStructValues(stackModel, excludeList...)
			formattedTimeValues = append(formattedTimeValues, originalValue[i])
		}
	}

	return formattedTimeValues, nil
}
