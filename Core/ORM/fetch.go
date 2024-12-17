package ORM

import (
	"fmt"
	"reflect"

	"github.com/superg3m/stoic-go/Core/Utility"
)

func Fetch[T InterfaceCRUD](sql string, bindParams ...any) *T {
	var dest T

	v := reflect.ValueOf(&dest).Elem()
	Utility.AssertMsg(v.Kind() == reflect.Struct, "Fetch: type %T is not a struct", dest)
	row := GetInstance().QueryRowx(sql, bindParams...)

	err := row.StructScan(&dest)
	Utility.AssertOnErrorMsg(err, "Fetch: failed to scan row into map: %s", err)

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
