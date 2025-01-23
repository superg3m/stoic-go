package UserToken

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
	"time"
)

type UserToken struct {
	DB      *sqlx.DB
	ID      int
	UserID  int
	Created time.Time
	Context string
	Token   string
}

func New() *UserToken {
	ret := new(UserToken)

	//ret.DB = ORM.GetInstance()
	ret.ID = 0
	ret.UserID = 0
	ret.Created = time.Now()
	ret.Context = ""
	ret.Token = ""

	return ret
}

func FromID(ID int) (*UserToken, error) {
	ret := New()
	ret.ID = ID
	read := ret.Read()
	if read.IsBad() {
		return nil, read.GetError()
	}

	return ret, nil
}

func init() {
	ORM.RegisterTableName(&UserToken{})
	ORM.RegisterTableColumn("ID", "ID", ORM.AUTO_INCREMENT|ORM.KEY)
	ORM.RegisterTableColumn("UserID", "UserID")
	ORM.RegisterTableColumn("Created", "Created")
	ORM.RegisterTableColumn("Context", "Context")
	ORM.RegisterTableColumn("Token", "Token")
}
