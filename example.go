package CLS

import (
	"errors"
	"time"
)

var ( // ERRORS
	ERROR_INVALID_EMAIL = errors.New("Invalid Email")
	ERROR_INVALID_ID    = errors.New("Invalid ID")
)

type User struct {
	BaseStoicTable

	ID             int
	Email          string
	EmailConfirmed bool
	Joined         time.Time
	LastActive     time.Time
	LastLogin      time.Time
}

func FromID(id int) {

}

func FromEmail(email string) (User, error) {
	var ret User

	if !Utils.ValidEmail() {
		return User{}, ERROR_INVALID_EMAIL
	}
}

// Implement the setupTable function for User
func init() {
	ORM.RegisterTableName("User")
	ORM.RegisterTableColumn("ID", "user_id", PRIMARY_KEY) // Using reflection to know the type!
	ORM.RegisterTableColumn("Email", "email_address", PRIMARY_KEY)
	ORM.RegisterTableColumn("Name", "full_name", NULLABLE|UPDATABLE)
	ORM.RegisterTableColumn("Age", "age", PRIMARY_KEY)
	ORM.RegisterTableColumn("ID", "user_id", PRIMARY_KEY)
}
