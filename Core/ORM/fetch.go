package ORM

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

// Fetch maps database row to the destination struct (dest).
func Fetch[T any](row *sqlx.Row) (T, error) {
	var dest T

	v := reflect.ValueOf(&dest).Elem()
	if v.Kind() != reflect.Struct {
		return dest, fmt.Errorf("Fetch: type %T is not a struct", dest)
	}

	if err := row.StructScan(&dest); err != nil {
		return dest, fmt.Errorf("Fetch: failed to scan row into struct: %w", err)
	}

	return dest, nil
}

func FetchAll[T any](rows *sqlx.Rows) ([]T, error) {
	var results []T

	for rows.Next() {
		var dest T

		if err := rows.StructScan(&dest); err != nil {
			return nil, fmt.Errorf("FetchAll: failed to scan row into struct: %w", err)
		}

		results = append(results, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchAll: error during rows iteration: %w", err)
	}

	return results, nil
}
