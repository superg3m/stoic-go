package User

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
	"time"
)


type User struct {
	DB *sqlx.DB
	ID int
	Email string
	EmailConfirmed bool
	Joined time.Time
	LastLogin *time.Time
	LastActive *time.Time
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

func FromID(ID int) (*User, []string) {
    ret := New()
    ret.ID = ID
    read := ret.Read()
    if read.IsBad() {
        return nil, read.GetErrors()
    }

    return ret, nil
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