package LoginKey

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
)

type LoginKeyProvider int

const (
	PASSWORD LoginKeyProvider = iota
	FACEBOOK
	TWITTER
	TWITCH
	GITHUB
	REDDIT
)

func isValidProvider(value int) bool {
	return value >= int(PASSWORD) && value <= int(REDDIT)
}

func getProvider(value int) LoginKeyProvider {
	Utility.AssertMsg(isValidProvider(value), "Invalid Provider value: %d", value)

	return LoginKeyProvider(value)
}

type LoginKey struct {
	DB *sqlx.DB

	UserID   int              `db:"UserID,   KEY"`
	Provider LoginKeyProvider `db:"Provider, KEY"`
	Key      string           `db:"Key,      UPDATABLE"`
}

func New() *LoginKey {
	ret := new(LoginKey)

	//ret.DB = ORM.GetInstance()
	ret.UserID = 0
	ret.Provider = PASSWORD
	ret.Key = ""

	return ret
}

func (loginKey *LoginKey) HashKey() {
	loginKey.Key = Utility.Sha256HashString(loginKey.Key)
}

func FromUserID_Provider(UserID int, Provider LoginKeyProvider) *LoginKey {
	ret := New()
	ret.UserID = UserID
	ret.Provider = getProvider(int(Provider))

	read := ret.Read()
	if read.IsBad() {
		return nil
	}

	ret.SetCache()

	return ret
}

func init() {
	ORM.RegisterModel(&LoginKey{})
}
