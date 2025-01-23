package UserSettings

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

type UserSettings struct {
	DB         *sqlx.DB
	UserID     int
	HtmlEmails bool
	PlaySounds bool
}

func New() *UserSettings {
	ret := new(UserSettings)

	//ret.DB = ORM.GetInstance()
	ret.UserID = 0
	ret.HtmlEmails = false
	ret.PlaySounds = false

	return ret
}

func FromUserID(UserID int) (*UserSettings, error) {
	ret := New()
	ret.UserID = UserID
	read := ret.Read()
	if read.IsBad() {
		return nil, read.GetError()
	}

	return ret, nil
}

func init() {
	ORM.RegisterTableName(&UserSettings{})
	ORM.RegisterTableColumn("UserID", "UserID", ORM.KEY)
	ORM.RegisterTableColumn("HtmlEmails", "HtmlEmails")
	ORM.RegisterTableColumn("PlaySounds", "PlaySounds")
}
