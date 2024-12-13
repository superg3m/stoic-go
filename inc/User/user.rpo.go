package User

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
)

func AllFromEmail(email string) ([]*User, error) {
	if !Utility.ValidEmail(email) {
		return nil, ERROR_INVALID_EMAIL
	}

	query := "SELECT * FROM User WHERE email = ?"
	rows, err := ORM.GetInstance().Queryx(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, _ := ORM.FetchAll[User](rows)

	return users, nil
}
