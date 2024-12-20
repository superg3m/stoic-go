package LoginKey

import (
	"errors"
	"github.com/superg3m/stoic-go/Core/Utility"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

type LoginKeyProvider int

const (
	ERROR LoginKeyProvider = iota
	PASSWORD
	FACEBOOK
	TWITTER
	TWITCH
	GITHUB
	REDDIT
)

type LoginKey struct {
	DB *sqlx.DB

	UserID   int
	Key      string
	Provider LoginKeyProvider
}

func New() *LoginKey {
	loginKey := new(LoginKey)

	loginKey.DB = ORM.GetInstance()

	loginKey.UserID = 0
	loginKey.Key = ""
	loginKey.Provider = ERROR

	return loginKey
}

func FromID(id int) *LoginKey {
	/*
		if id <= 0 {
			return nil
		}
		// sql := "SELECT * FROM User WHERE id = ?"
		// ORM.Fetch[User](sql, id)
	*/

	loginKey := New()
	loginKey.UserID = id
	read := loginKey.Read()

	if read.IsBad() {
		return nil
	}

	LoginKey.SetCache(*user)

	return user
}

func FromEmail(email string) *User {
	/*
		if !Utility.ValidEmail(email) {
			return nil
		}
		// sql := "SELECT * FROM User WHERE email = ?"
		// ORM.Fetch[User](sql, email)
	*/

	user := New()
	user.Email = email
	read := user.Read()

	if read.IsBad() {
		return nil
	}

	User.SetCache(*user)

	return user
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName(User{})
	ORM.RegisterTableColumn("ID", "ID", ORM.KEY, ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Username", "Username", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Password", "Password", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Email", "Email", ORM.UPDATABLE, ORM.UNIQUE)
	ORM.RegisterTableColumn("EmailConfirmed", "EmailConfirmed", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Joined", "Joined")
	ORM.RegisterTableColumn("LastLogin", "LastLogin", ORM.UPDATABLE)
	ORM.RegisterTableColumn("LastActive", "LastActive", ORM.UPDATABLE)
}
