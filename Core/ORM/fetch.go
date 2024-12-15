package ORM

import (
	"fmt"
	"reflect"

	"github.com/superg3m/stoic-go/Core/Utility"
)

func manuallyGetLastInsertID(tableName, columnName string) int64 {
	var dest int64

	v := reflect.ValueOf(&dest).Elem()
	Utility.AssertMsg(v.Kind() != reflect.Struct, fmt.Sprintf("Fetch: type %T is not a struct", dest))

	sql := `
	SELECT IFNULL(MAX(%s), 0)
	FROM %s
	LIMIT 1;
	`
	sql = fmt.Sprintf(sql, columnName, tableName)

	row := GetInstance().QueryRowx(sql)

	err := row.Scan(&dest)
	fmt.Println(dest)
	Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row: %s", err))

	return dest
}

// Fetch maps database row to the destination struct (dest).
func Fetch[T any](sql string, bindParams ...any) *T {
	var dest T

	v := reflect.ValueOf(&dest).Elem()
	Utility.AssertMsg(v.Kind() != reflect.Struct, fmt.Sprintf("Fetch: type %T is not a struct", dest))

	row := GetInstance().QueryRowx(sql, bindParams...)

	err := row.StructScan(&dest)
	Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row into struct: %s", err))

	return &dest
}

func FetchAll[T InterfaceCRUD](sql string, bindParams ...any) ([]*T, error) {
	var results []*T

	rows, errQuery := GetInstance().Queryx(sql, bindParams...)
	if errQuery != nil {
		return nil, errQuery
	}

	defer rows.Close()

	for rows.Next() {
		var dest T

		err := rows.StructScan(&dest)
		Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row into struct: %s", err))

		results = append(results, &dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchAll: error during rows iteration: %s", err)
	}

	return results, nil
}
