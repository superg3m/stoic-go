package User

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

var (
	ERROR_INVALID_EMAIL = errors.New("Invalid Email")
	ERROR_INVALID_ID    = errors.New("Invalid ID")
)

type User struct {
	DB *sqlx.DB

	ID             int    // Primary Key
	Email          string // Unique
	EmailConfirmed bool
	Joined         time.Time
	LastLogin      *time.Time
	LastActive     *time.Time
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

func FromID(id int) (*User, []string) {
	user := New()
	user.ID = id
	read := user.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	user.SetCache()

	return user, nil
}

func FromEmail(email string) (*User, []string) {
	user := New()
	user.Email = email
	read := user.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	user.SetCache()

	return user, nil
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName(&User{})
	ORM.RegisterTableColumn("ID", "ID", ORM.KEY, ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Email", "Email", ORM.UPDATABLE, ORM.UNIQUE)
	ORM.RegisterTableColumn("EmailConfirmed", "EmailConfirmed", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Joined", "Joined")
	ORM.RegisterTableColumn("LastLogin", "LastLogin", ORM.UPDATABLE)
	ORM.RegisterTableColumn("LastActive", "LastActive", ORM.UPDATABLE)
}
