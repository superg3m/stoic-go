package User

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ERROR_INVALID_EMAIL = errors.New("Invalid Email")
	ERROR_INVALID_ID    = errors.New("Invalid ID")
)

type User struct {
	DB *sqlx.DB

	ID             int        `db:"ID,             KEY, AUTO_INCREMENT"`
	Email          string     `db:"Email,          UNIQUE,UPDATABLE"`
	EmailConfirmed bool       `db:"EmailConfirmed, UPDATABLE"`
	Joined         time.Time  `db:"Joined"`
	LastLogin      *time.Time `db:"LastLogin,      UPDATABLE"`
	LastActive     *time.Time `db:"LastActive,     UPDATABLE"`
}

func New() *User {
	user := new(User)

	//user.DB = db

	user.ID = 0
	user.Email = ""
	user.EmailConfirmed = false
	user.Joined = time.Now()
	user.LastActive = nil
	user.LastLogin = nil

	return user
}

func FromID(id int) (*User, error) {
	user := New()
	user.ID = id
	read := user.Read()
	if read.IsBad() {
		return nil, read.GetError()
	}

	return user, nil
}

func FromEmail(email string) (*User, error) {
	user := New()
	user.Email = email
	read := user.Read()
	if read.IsBad() {
		return nil, read.GetError()
	}

	return user, nil
}
