package Api

import "time" // Adjust imports as necessary

type User struct {
	ORM.BaseStoicTable
	ID int // Description for ID
	Email string // Description for Email
	Joined time.Time // Description for Joined
}

// Implement the setupTable function for User
func init() {
	ORM.RegisterTableName("User")
	ORM.RegisterTableColumn("ID", "user_id", ORM.PRIMARY_KEY)
	ORM.RegisterTableColumn("Email", "email_address", ORM.NULLABLE|ORM.UPDATABLE)
	ORM.RegisterTableColumn("Joined", "joined_at", ORM.NULLABLE)
}