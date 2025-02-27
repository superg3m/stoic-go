package User

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
	"time"
)

const (
	ERROR_INVALID_EMAIL = "Invalid Email"
	ERROR_INVALID_ID    = "Invalid ID"
)

type User struct {
	Meta
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

func (u *User) CanCreate() []string {
	dbUser := *u
	read := dbUser.Read()
	if read.IsBad() {
		return nil
	}

	var errors []string = nil
	if !Utility.ValidEmail(u.Email) {
		errors = append(errors, "User Invalid Email")
	}

	if u.Email == dbUser.Email {
		errors = append(errors, "User Duplicate Email")
	}

	return errors
}

func (u *User) CanRead() []string {
	return nil
}

func (u *User) CanUpdate() []string {
	dbUser := *u
	read := dbUser.Read()
	if read.IsBad() {
		return nil
	}

	var errors []string = nil
	if !Utility.ValidEmail(u.Email) {
		errors = append(errors, "User Invalid Email")
	}

	if u.ID == dbUser.ID {
		return nil
	}

	if u.Email == dbUser.Email {
		errors = append(errors, "User Duplicate Email")
	}

	return errors
}

func (u *User) CanDelete() []string {
	return nil
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName(&User{})
	ORM.RegisterTableColumn("ID", "ID", ORM.KEY|ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Email", "Email", ORM.UPDATABLE|ORM.UNIQUE)
	ORM.RegisterTableColumn("EmailConfirmed", "EmailConfirmed", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Joined", "Joined")
	ORM.RegisterTableColumn("LastLogin", "LastLogin", ORM.UPDATABLE)
	ORM.RegisterTableColumn("LastActive", "LastActive", ORM.UPDATABLE)
}
