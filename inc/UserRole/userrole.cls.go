package UserRole

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

type UserRole struct {
	DB     *sqlx.DB
	UserID int
	RoleID int
}

func New() *UserRole {
	ret := new(UserRole)

	//ret.DB = ORM.GetInstance()
	ret.UserID = 0
	ret.RoleID = 0

	return ret
}

func FromUserID_RoleID(UserID int, RoleID int) (*UserRole, []string) {
	ret := New()
	ret.UserID = UserID
	ret.RoleID = RoleID
	read := ret.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	return ret, nil
}

func init() {
	ORM.RegisterTableName(&UserRole{})
	ORM.RegisterTableColumn("UserID", "UserID", ORM.KEY)
	ORM.RegisterTableColumn("RoleID", "RoleID", ORM.KEY)
}
