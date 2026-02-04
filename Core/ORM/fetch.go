package ORM

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
)

func Fetch[T InterfaceCRUD](db *sqlx.DB, sql string, bindParams ...any) (T, error) {
	dest := Utility.NewUnderlyingType[T]()

	err := db.Get(dest, sql, bindParams...)
	if err != nil {
		dest.SetCache()
	}

	return dest, err
}

func FetchAll[T InterfaceCRUD](db *sqlx.DB, sql string, bindParams ...any) ([]T, error) {
	var results []T

	err := db.Select(&results, sql, bindParams...)
	if err != nil {
		for _, result := range results {
			result.SetCache()
		}
	}

	return results, err
}
