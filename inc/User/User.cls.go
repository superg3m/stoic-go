package User

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/superg3m/stoic-go/Core/Utility"

	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

var (
	ERROR_INVALID_EMAIL = errors.New("Invalid Email")
	ERROR_INVALID_ID    = errors.New("Invalid ID")
)

type User struct {
	DB *sqlx.DB `json:"-"`

	ID             int        `json:"ID"`
	Email          string     `json:"Email"`
	EmailConfirmed bool       `json:"EmailConfirmed"`
	Joined         time.Time  `json:"Joined"`
	LastLogin      *time.Time `json:"LastLogin"`
	LastActive     *time.Time `json:"LastActive"`
}

type CookieData struct {
	ID int `json:"ID"`
}

var DatabaseName = "stoic"

func New() *User {
	user := new(User)

	user.DB = ORM.GetInstance(DatabaseName)

	user.ID = 0
	user.Email = ""
	user.EmailConfirmed = false
	user.Joined = time.Now()
	user.LastActive = nil
	user.LastLogin = nil

	return user
}

func FromID(id int) (*User, []string) {
	user := New()
	user.ID = id
	read := user.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	user.SetCache()

	return user, nil
}

func FromEmail(email string) (*User, []string) {
	user := New()
	user.Email = email
	read := user.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	user.SetCache()

	return user, nil
}

func AllFromEmail(email string) ([]*User, error) {
	if !Utility.ValidEmail(email) {
		return nil, ERROR_INVALID_EMAIL
	}

	sql := "SELECT * FROM User WHERE email = ?"
	users, _ := ORM.FetchAll[*User](ORM.GetInstance(DatabaseName), sql, email)

	return users, nil
}

func GetUserList() (string, error) {
	var admins []*User

	sql := `
	SELECT * 
	FROM User LEFT JOIN UserRole ON User.ID = UserRole.UserID
	WHERE UserRole.ID IS NOT NULL
	`

	admins, err := ORM.FetchAll[*User](ORM.GetInstance(DatabaseName), sql)
	if err != nil {
		fmt.Printf("Error getting admins: %s\n", err)
		return "", fmt.Errorf("getExcludedUserList: %w", err)
	}

	var ret []string

	for _, user := range admins {
		ret = append(ret, user.Email)
	}

	return strings.Join(ret, ","), nil
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName(&User{})
	ORM.RegisterTableColumn("ID", "ID", ORM.KEY, ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Email", "Email", ORM.UPDATABLE, ORM.UNIQUE)
	ORM.RegisterTableColumn("EmailConfirmed", "EmailConfirmed", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Joined", "Joined")
	ORM.RegisterTableColumn("LastLogin", "LastLogin", ORM.UPDATABLE)
	ORM.RegisterTableColumn("LastActive", "LastActive", ORM.UPDATABLE)
}
