package ORM

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"
)

func Fetch[T InterfaceCRUD](db *sqlx.DB, sql string, bindParams ...any) (T, error) {
	var dest T

	tType := reflect.TypeOf(dest)
	if tType.Kind() == reflect.Ptr {
		dest = reflect.New(tType.Elem()).Interface().(T)
	} else {
		dest = *new(T)
	}

	row := db.QueryRowx(sql, bindParams...)

	pointers := Utility.GetStructMemberPointer(dest, excludeList...)
	err := row.Scan(pointers...)
	if err != nil {
		return dest, err
	}

	dest.SetCache()

	return dest, nil
}

func FetchAll[T InterfaceCRUD](db *sqlx.DB, sql string, bindParams ...any) ([]T, error) {
	var results []T

	rows, errQuery := db.Queryx(sql, bindParams...)
	if errQuery != nil {
		return nil, errQuery
	}

	defer rows.Close()

	for rows.Next() {
		newValuePtr := reflect.New(reflect.TypeOf((*T)(nil)).Elem().Elem())
		newValue := newValuePtr.Interface().(T)

		pointers := Utility.GetStructMemberPointer(newValue, excludeList...)
		if err := rows.Scan(pointers...); err != nil {
			return nil, fmt.Errorf("Fetch: failed to scan row into struct: %w", err)
		}

		results = append(results, newValue)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchAll: row iteration error: %w", err)
	}

	return results, nil
}
