package Api

import (
	"errors"
	"github.com/superg3m/stoic-go/cmd/src/ORM"
	"github.com/superg3m/stoic-go/cmd/src/Utility"
	"time"
)

var ( // ERRORS
	ERROR_INVALID_EMAIL = errors.New("Invalid Email")
	ERROR_INVALID_ID    = errors.New("Invalid ID")
)

type User struct {
	ORM.BaseStoicTable

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

	if !Utility.ValidEmail(email) {
		return User{}, ERROR_INVALID_EMAIL
	}

	return ret, nil
}

// Implement the setupTable function for User
func init() {
	ORM.RegisterTableName("User")
	ORM.RegisterTableColumn("ID", "user_id", ORM.PRIMARY_KEY) // Using reflection to know the type!
	ORM.RegisterTableColumn("Email", "email_address", ORM.PRIMARY_KEY)
	ORM.RegisterTableColumn("Name", "full_name", ORM.NULLABLE|ORM.UPDATABLE)
	ORM.RegisterTableColumn("Age", "age", ORM.PRIMARY_KEY)
	ORM.RegisterTableColumn("ID", "user_id", ORM.PRIMARY_KEY)
}
