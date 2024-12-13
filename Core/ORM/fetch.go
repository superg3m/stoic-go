package ORM

import (
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
)

// Fetch maps database row to the destination struct (dest).
func Fetch[T InterfaceCRUD](row *sqlx.Row) *T {
	var dest T

	v := reflect.ValueOf(&dest).Elem()
	Utility.AssertMsg(v.Kind() != reflect.Struct, fmt.Sprintf("Fetch: type %T is not a struct", dest))

	err := row.StructScan(&dest)
	Utility.AssertOnErrorMsg(err, fmt.Sprintf("Fetch: failed to scan row into struct: %s", err))

	return &dest
}

func FetchAll[T InterfaceCRUD](rows *sqlx.Rows) ([]*T, error) {
	var results []*T

	for rows.Next() {
		var dest T

		if err := rows.StructScan(&dest); err != nil {
			return nil, fmt.Errorf("FetchAll: failed to scan row into struct: %s", err)
		}

		results = append(results, &dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchAll: error during rows iteration: %s", err)
	}

	return results, nil
}
