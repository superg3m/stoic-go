package UserVisibilities

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

type UserVisibilities struct {
	DB          *sqlx.DB `json:"-"`
	UserID      int
	RealName    string
	Description string
	Gender      string
}

func New() *UserVisibilities {
	ret := new(UserVisibilities)

	//ret.DB = ORM.GetInstance()
	ret.UserID = 0
	ret.RealName = ""
	ret.Description = ""
	ret.Gender = ""

	return ret
}

func FromUserID(UserID int) (*UserVisibilities, []string) {
	ret := New()
	ret.UserID = UserID
	read := ret.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	return ret, nil
}

func init() {
	ORM.RegisterTableName(&UserVisibilities{})
	ORM.RegisterTableColumn("UserID", "UserID", ORM.KEY)
	ORM.RegisterTableColumn("RealName", "RealName")
	ORM.RegisterTableColumn("Description", "Description")
	ORM.RegisterTableColumn("Gender", "Gender")
}
