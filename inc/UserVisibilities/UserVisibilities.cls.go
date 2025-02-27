package UserVisibilities

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

type UserVisibilities struct {
	DB          *sqlx.DB
	UserID      int
	Profile     bool
	Email       bool
	Searches    bool
	Birthday    bool
	RealName    bool
	Description bool
	Gender      bool
}

func New() *UserVisibilities {
	ret := new(UserVisibilities)

	//ret.DB = ORM.GetInstance()
	ret.UserID = 0
	ret.Profile = false
	ret.Email = false
	ret.Searches = false
	ret.Birthday = false
	ret.RealName = false
	ret.Description = false
	ret.Gender = false

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
	ORM.RegisterTableColumn("Profile", "Profile")
	ORM.RegisterTableColumn("Email", "Email")
	ORM.RegisterTableColumn("Searches", "Searches")
	ORM.RegisterTableColumn("Birthday", "Birthday")
	ORM.RegisterTableColumn("RealName", "RealName")
	ORM.RegisterTableColumn("Description", "Description")
	ORM.RegisterTableColumn("Gender", "Gender")
}
