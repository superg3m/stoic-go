package LoginKey

import (
	"github.com/jmoiron/sqlx"
)

type Meta struct {
	DB       *sqlx.DB
	UserID   int
	Provider int
	Key      string
}

func FromUserID_Provider(UserID int, Provider int) (*LoginKey, []string) {
	ret := New()
	ret.UserID = UserID
	ret.Provider = Provider

	read := ret.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	return ret, nil
}
