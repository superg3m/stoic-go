package User

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
	"time"
)


type User struct {
	DB *sqlx.DB
	ID int
	Username string
	Joined time.Time
	Email string
}

func New() *User {
	ret := new(User)

	//ret.DB = ORM.GetInstance()
    ret.ID = 0
    ret.Username = ""
    ret.Joined = time.Now()
    ret.Email = ""

	return ret
}

func FromID_Username(ID int, Username string) *User {
    ret := New()
    ret.ID = ID
    ret.Username = Username

    ret.Read()

    return ret
}
func FromEmail(Email string) *User {
    ret := New()
    ret.Email = Email

    ret.Read()

    return ret
}
func FromJoined(Joined time.Time) *User {
    ret := New()
    ret.Joined = Joined

    ret.Read()

    return ret
}

func init() {
	ORM.RegisterTableName(User{})
	ORM.RegisterTableColumn("ID", "user_id", ORM.KEY)
	ORM.RegisterTableColumn("Username", "userName", ORM.KEY)
	ORM.RegisterTableColumn("Joined", "joined", ORM.UNIQUE)
	ORM.RegisterTableColumn("Email", "email_address", ORM.UNIQUE|ORM.UPDATABLE)
}