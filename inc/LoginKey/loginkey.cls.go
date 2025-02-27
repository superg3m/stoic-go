package LoginKey

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
)

const (
	PASSWORD int = iota
	FACEBOOK
	TWITTER
	TWITCH
	GITHUB
	REDDIT
)

func isValidProvider(value int) bool {
	return value >= PASSWORD && value <= REDDIT
}

type LoginKey struct {
	Meta
}

func New() *LoginKey {
	ret := new(LoginKey)

	//ret.DB = ORM.GetInstance()
	ret.UserID = 0
	ret.Provider = 0
	ret.Key = ""

	return ret
}

func (model *LoginKey) CanCreate() []string {
	return nil
}

func (model *LoginKey) CanRead() []string {
	return nil
}

func (model *LoginKey) CanUpdate() []string {
	return nil
}

func (model *LoginKey) CanDelete() []string {
	return nil
}

func (loginKey *LoginKey) HashKey() {
	loginKey.Key = Utility.Sha256HashString(loginKey.Key)
}

func init() {
	ORM.RegisterTableName(&LoginKey{})
	ORM.RegisterTableColumn("UserID", "UserID", ORM.KEY)
	ORM.RegisterTableColumn("Provider", "Provider", ORM.KEY)
	ORM.RegisterTableColumn("Key", "Key", ORM.UPDATABLE)
}
