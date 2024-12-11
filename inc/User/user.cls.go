package User

import (
	"errors"
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
)

var (
	ERROR_INVALID_EMAIL = errors.New("Invalid Email")
	ERROR_INVALID_ID    = errors.New("Invalid ID")
)

type User struct {
	ORM.StoicModel

	ID             int
	Username       string
	Password       string
	Email          string
	EmailConfirmed bool
	// Joined         time.Time
	// LastActive     time.Time
	// LastLogin      time.Time
}

func (u *User) CanCreate() bool {
	return true
}

func (u *User) CanUpdate() bool {
	return true
}

func (u *User) CanDelete() bool {
	return true
}

var _ ORM.InterfaceCRUD = (*User)(nil)

func New() *User {
	user := new(User)
	user.DB = ORM.GetInstance()
	user.TableName = "User"

	return user
}

func FromID(id int) (*User, error) {
	var user User

	if id <= 0 {
		return nil, ERROR_INVALID_ID
	}

	query := "SELECT * FROM User WHERE id = ?"
	row := ORM.GetInstance().QueryRowx(query, id)
	user = ORM.Fetch[User](row)

	return &user, nil
}

func FromEmail(email string) (*User, error) {
	if !Utility.ValidEmail(email) {
		return nil, ERROR_INVALID_EMAIL
	}

	query := "SELECT * FROM User WHERE email = ?"
	row := ORM.GetInstance().QueryRowx(query, email)
	user := ORM.Fetch[User](row)

	return &user, nil
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName("User")
	ORM.RegisterTableColumn("ID", "id", ORM.PRIMARY_KEY)
	ORM.RegisterTableColumn("Email", "email", ORM.UNIQUE)
	ORM.RegisterTableColumn("EmailConfirmed", "email_confirmed", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Username", "username", ORM.NULLABLE|ORM.UPDATABLE)
	ORM.RegisterTableColumn("Password", "password", ORM.UPDATABLE)
	// ORM.RegisterTableColumn("Joined", "joined", ORM.NULLABLE)
	// ORM.RegisterTableColumn("LastActive", "last_active", ORM.NULLABLE)
	// ORM.RegisterTableColumn("LastLogin", "last_login", ORM.NULLABLE)
}
