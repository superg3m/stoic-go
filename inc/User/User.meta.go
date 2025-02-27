package User

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Meta struct {
	DB *sqlx.DB
	ID int
	Email string
	EmailConfirmed bool
	Joined time.Time
	LastLogin *time.Time
	LastActive *time.Time
}

func FromID(ID int) (*User, []string) {
    ret := New()
    ret.ID = ID
    read := ret.Read()
    if read.IsBad() {
        return nil, read.GetErrors()
    }

    return ret, nil
}