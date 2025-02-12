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
	DB       *sqlx.DB
	UserID   int
	Provider LoginKeyProvider
	Key      string
}

func New() *LoginKey {
	ret := new(LoginKey)

	//ret.DB = ORM.GetInstance()
	ret.UserID = 0
	ret.Provider = 0
	ret.Key = ""

	return ret
}

func (loginKey *LoginKey) HashKey() {
	loginKey.Key = Utility.Sha256HashString(loginKey.Key)
}

func FromUserID_Provider(UserID int, Provider LoginKeyProvider) (*LoginKey, []string) {
	ret := New()
	ret.UserID = UserID
	ret.Provider = getProvider(int(Provider))

	read := ret.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	return ret, nil
}

func init() {
	ORM.RegisterTableName(&LoginKey{})
	ORM.RegisterTableColumn("UserID", "UserID", ORM.KEY)
	ORM.RegisterTableColumn("Provider", "Provider", ORM.KEY)
	ORM.RegisterTableColumn("Key", "Key", ORM.UPDATABLE)
}
