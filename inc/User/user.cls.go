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
	ID             int // Not updatable
	Username       string
	Password       string
	Email          string
	EmailConfirmed bool
	// Joined         time.Time
	// LastActive     time.Time
	// LastLogin      time.Time
}

func New() *User {
	user := new(User)

	user.ID = 0
	user.Username = ""
	user.Password = ""
	user.Email = ""
	user.EmailConfirmed = false

	return user
}

func FromID(id int) *User {
	if id <= 0 {
		return nil
	}

	sql := "SELECT * FROM User WHERE id = ?"

	return ORM.Fetch[User](sql, id)
}

func FromEmail(email string) *User {
	if !Utility.ValidEmail(email) {
		return nil
	}

	sql := "SELECT * FROM User WHERE email = ?"

	return ORM.Fetch[User](sql, email)
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName("User")
	ORM.RegisterTableColumn("ID", "id", ORM.KEY, ORM.AUTO_INCREMENT)
	ORM.RegisterTableColumn("Username", "username", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Password", "password", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Email", "email", ORM.UPDATABLE)
	ORM.RegisterTableColumn("EmailConfirmed", "email_confirmed", ORM.UPDATABLE)

	// ORM.RegisterTableColumn("Joined", "joined", ORM.NULLABLE)
	// ORM.RegisterTableColumn("LastActive", "last_active", ORM.NULLABLE)
	// ORM.RegisterTableColumn("LastLogin", "last_login", ORM.NULLABLE)
}
