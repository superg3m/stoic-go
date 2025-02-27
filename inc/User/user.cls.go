package User

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"time"
)

const ERROR_INVALID_EMAIL = "Invalid Email"

type User struct {
	Meta
}

func New() *User {
	ret := new(User)

	//ret.DB = ORM.GetInstance()
	ret.ID = 0
	ret.Email = ""
	ret.EmailConfirmed = false
	ret.Joined = time.Now()
	ret.LastLogin = nil
	ret.LastActive = nil

	return ret
}

func (model *User) CanCreate() []string {
	return nil
}

func (model *User) CanRead() []string {
	return nil
}

func (model *User) CanUpdate() []string {
	return nil
}

func (model *User) CanDelete() []string {
	return nil
}

func init() {
	ORM.RegisterTableName(&User{})
	ORM.RegisterTableColumn("ID", "ID", ORM.KEY|ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Email", "Email", ORM.UNIQUE)
	ORM.RegisterTableColumn("EmailConfirmed", "EmailConfirmed")
	ORM.RegisterTableColumn("Joined", "Joined")
	ORM.RegisterTableColumn("LastLogin", "LastLogin", ORM.NULLABLE)
	ORM.RegisterTableColumn("LastActive", "LastActive", ORM.NULLABLE)
}
