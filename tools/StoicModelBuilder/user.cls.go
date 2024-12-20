package User

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
	"time"
)


type User struct {
	DB *sqlx.DB
	ID int
	Username any
	Email any
	Password any
	EmailConfirmed bool
	Joined *time.Time
	LastLogin *time.Time
	LastActive *time.Time
}

func New() *User {
	ret := new(User)

	//ret.DB = ORM.GetInstance()
    ret.ID = 0
    ret.EmailConfirmed = false
    ret.Joined = nil
    ret.LastLogin = nil
    ret.LastActive = nil

	return ret
}

func From() *User {
    ret := New()

    ret.Read()

    return ret
}

func init() {
	ORM.RegisterTableName(User{})
	ORM.RegisterTableColumn("ID", "ID", ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Username", "Username", )
	ORM.RegisterTableColumn("Email", "Email", )
	ORM.RegisterTableColumn("Password", "Password", )
	ORM.RegisterTableColumn("EmailConfirmed", "EmailConfirmed", )
	ORM.RegisterTableColumn("Joined", "Joined", )
	ORM.RegisterTableColumn("LastLogin", "LastLogin", )
	ORM.RegisterTableColumn("LastActive", "LastActive", )
}