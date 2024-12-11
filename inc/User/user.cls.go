package User

import (
	"errors"
	"github.com/superg3m/stoic-go/Core/Database"
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
	"time"
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
	Joined         time.Time
	LastActive     time.Time
	LastLogin      time.Time
}

func FromID(id int) (*User, error) {
	var user User

	if id <= 0 {
		return nil, ERROR_INVALID_ID
	}

	query := "SELECT * FROM User WHERE id = ?"
	err := user.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.EmailConfirmed, &user.Joined, &user.LastActive, &user.LastLogin)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FromEmail(email string) (*User, error) {
	if !Utility.ValidEmail(email) {
		return nil, ERROR_INVALID_EMAIL
	}

	query := "SELECT * FROM User WHERE email = ?"
	row := Database.GetInstance().QueryRowx(query, email)
	user, err := ORM.Fetch[User](row)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Register ORM metadata
func init() {
	ORM.RegisterTableName("User")
	ORM.RegisterTableColumn("ID", "user_id", ORM.PRIMARY_KEY)
	ORM.RegisterTableColumn("Email", "email_address", ORM.UNIQUE)
	ORM.RegisterTableColumn("Username", "username", ORM.NULLABLE|ORM.UPDATABLE)
	ORM.RegisterTableColumn("Joined", "joined", ORM.NULLABLE)
	ORM.RegisterTableColumn("LastActive", "last_active", ORM.NULLABLE)
	ORM.RegisterTableColumn("LastLogin", "last_login", ORM.NULLABLE)
}
