package ORM

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
)

// mysql
// sqlserver
// postgres
// sql_lite

func CreateRecord(db *sqlx.DB, payload ModelPayload) (sql.Result, error) {
	placeholders := make([]string, len(payload.ColumnNames))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	var newDbNames []string
	var newPlaceholders []string
	var newValues []any

	for i, fieldName := range payload.MemberNames {
		attribute := globalTable[payload.TableName][fieldName]
		if attribute.isAutoIncrement() {
			continue
		}
		newDbNames = append(newDbNames, "`"+payload.ColumnNames[i]+"`")
		newPlaceholders = append(newPlaceholders, placeholders[i])
		newValues = append(newValues, payload.Values[i])
	}

	query := fmt.Sprintf(
		"INSERT INTO `%s` (%s) VALUES (%s)",
		payload.TableName,
		strings.Join(newDbNames, ", "),
		strings.Join(newPlaceholders, ", "),
	)

	result, execErr := db.Exec(query, newValues...)
	if execErr != nil {
		return result, fmt.Errorf("failed to execute query: %s\nError: %w", query, execErr)
	}

	return result, nil
}

func ReadRecord[T InterfaceCRUD](db *sqlx.DB, payload ModelPayload, model T) error {
	pKeyQuery, uniqueQueries, _ := buildSQLReadQueries(payload)

	{
		allMatch := true
		for i, pointer := range payload.UniquePointers {
			query := uniqueQueries[i]
			temp, err := Fetch[T](query, pointer)
			if err != nil {
				allMatch = false
			}

			if i == len(payload.UniquePointers)-1 && allMatch {
				Utility.Copy(temp, model)
				return nil
			}
		}
	}

	// ----------------------------------------------------------

	{
		temp, err := Fetch[T](pKeyQuery, payload.PrimaryKeyPointers...)
		if err == nil {
			Utility.Copy(temp, model)
			return nil
		}
	}

	// ----------------------------------------------------------

	return errors.New("failed to fetch record")
}

func UpdateRecord(db *sqlx.DB, payload ModelPayload) (sql.Result, error) {
	keyColumns := getColumnNames(payload.TableName, payload.PrimaryKeyMemberNames)
	updateFields := payload.ColumnNames
	updateValues := payload.Values

	var updates []string
	for _, fieldName := range updateFields {
		updates = append(updates, fmt.Sprintf("`%s` = ?", fieldName))
	}

	var where []string
	for _, columnNames := range keyColumns {
		where = append(where, fmt.Sprintf("`%s` = ?", columnNames))
	}

	query := fmt.Sprintf(
		"UPDATE `%s` SET %s WHERE %s",
		payload.TableName,
		strings.Join(updates, ", "),
		strings.Join(where, " AND "),
	)

	result, execErr := db.Exec(query, append(updateValues, payload.PrimaryKeyPointers...)...)
	if execErr != nil {
		return result, fmt.Errorf("failed to execute query: %w", execErr)
	}

	return result, nil
}

func DeleteRecord(db *sqlx.DB, payload ModelPayload) (sql.Result, error) {
	var conditions []string
	primaryKeyColumns := getColumnNames(payload.TableName, payload.PrimaryKeyMemberNames)

	for _, fieldName := range primaryKeyColumns {
		conditions = append(conditions, fmt.Sprintf("`%s` = ?", fieldName))
	}

	query := fmt.Sprintf(
		"DELETE FROM `%s` WHERE %s",
		payload.TableName,
		strings.Join(conditions, " AND "),
	)

	result, execErr := db.Exec(query, payload.PrimaryKeyPointers...)
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

func buildSQLReadQueries(payload ModelPayload) (primaryQuery string, uniqueQueries []string, err error) {
	Utility.Assert(len(payload.ColumnNames) > 0)

	primaryKeyColumns := getColumnNames(payload.TableName, payload.PrimaryKeyMemberNames)
	if len(primaryKeyColumns) == 0 {
		return "", nil, fmt.Errorf("no primary keys defined for table: %s", payload.TableName)
	}

	primaryPlaceholders := make([]string, len(primaryKeyColumns))
	for i := range primaryPlaceholders {
		primaryPlaceholders[i] = "?"
	}
	primaryQuery = fmt.Sprintf(
		"SELECT * FROM %s WHERE %s",
		payload.TableName,
		strings.Join(primaryKeyColumns, " = ? AND ")+" = ?",
	)

	uniqueKeyGroups := getColumnNames(payload.TableName, payload.UniqueMemberNames)
	for _, uniqueKey := range uniqueKeyGroups {
		uniqueQuery := fmt.Sprintf(
			"SELECT * FROM %s WHERE %s = ?",
			payload.TableName,
			uniqueKey,
		)
		uniqueQueries = append(uniqueQueries, uniqueQuery)
	}

	return primaryQuery, uniqueQueries, nil
}
