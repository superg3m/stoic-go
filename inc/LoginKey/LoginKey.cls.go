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
	DB *sqlx.DB `json:"-"`

	UserID   int              `db:"UserID"   json:"UserID"`
	Provider LoginKeyProvider `db:"Provider" json:"Provider"`
	Key      string           `db:"Key"      json:"Key"`
}

var DatabaseName = "stoic"

func New() *LoginKey {
	ret := new(LoginKey)

	ret.DB = ORM.GetInstance(DatabaseName)
	ret.UserID = 0
	ret.Provider = 0
	ret.Key = ""

	return ret
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
