package Database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/core/Utility"
	"strings"
)

// mysql
// sqlserver
// postgres
// sql_lite

func InsertIntoDB[T any](db *sql.DB, tableName string, table T) error {
	fieldNames := Utility.GetStructFieldNames(table)

	placeholders := make([]string, len(fieldNames))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	values := Utility.GetStructValues(table)

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(fieldNames, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := db.Exec(query, values...)
	return err
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

func ConnectToDatabase(dbEngine, dsn string) *sqlx.DB {
	db, err := sqlx.Connect(dbEngine, dsn)
	Utility.AssertOnErrorMsg(err, "Failed to connect to database")
	return db
}
