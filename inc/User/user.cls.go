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

	ID             int
	Username       string
	Password       string
	Email          string
	EmailConfirmed bool
	Joined         time.Time
	LastLogin      *time.Time
	LastActive     *time.Time
}

func New() *User {
	user := new(User)

	//user.DB = db

	user.ID = 0
	user.Username = ""
	user.Password = ""
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

	user.SetCache()

	return user, nil
}

func FromEmail(email string) (*User, error) {
	user := New()
	user.Email = email
	read := user.Read()
	if read.IsBad() {
		return nil, read.GetError()
	}

	user.SetCache()

	return user, nil
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName(&User{})
	ORM.RegisterTableColumn("ID", "ID", ORM.KEY, ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Username", "Username", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Password", "Password", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Email", "Email", ORM.UPDATABLE, ORM.UNIQUE)
	ORM.RegisterTableColumn("EmailConfirmed", "EmailConfirmed", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Joined", "Joined")
	ORM.RegisterTableColumn("LastLogin", "LastLogin", ORM.UPDATABLE)
	ORM.RegisterTableColumn("LastActive", "LastActive", ORM.UPDATABLE)
}
