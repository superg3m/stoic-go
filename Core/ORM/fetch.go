package ORM

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"
)

func Fetch[T InterfaceCRUD](sql string, bindParams ...any) (T, error) {
	var dest T

	tType := reflect.TypeOf(dest)
	if tType.Kind() == reflect.Ptr {
		dest = reflect.New(tType.Elem()).Interface().(T)
	} else {
		dest = *new(T)
	}

	row := GetInstance().QueryRowx(sql, bindParams...)

	err := row.StructScan(dest)
	if err != nil {
		return dest, err
	}

	dest.SetCache()

	return dest, nil
}

func FetchAll[T InterfaceCRUD](sql string, bindParams ...any) ([]T, error) {
	var results []T

	rows, errQuery := GetInstance().Queryx(sql, bindParams...)
	if errQuery != nil {
		return nil, errQuery
	}

	defer rows.Close()

	for rows.Next() {
		dest := *new(T)

		err := rows.StructScan(dest)
		Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row into struct: %s", err))

		results = append(results, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchAll: error during rows iteration: %s", err)
	}

	return results, nil
}
