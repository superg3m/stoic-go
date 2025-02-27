package User

import (
	"errors"
	"fmt"
	"strings"

	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Utility"
)

func AllFromEmail(email string) ([]*User, error) {
	if !Utility.ValidEmail(email) {
		return nil, errors.New(ERROR_INVALID_EMAIL)
	}

	sql := "SELECT * FROM User WHERE email = ?"
	users, _ := ORM.FetchAll[*User](sql, email)

	return users, nil
}

func GetUserList() (string, error) {
	var admins []*User

	sql := `
	SELECT * 
	FROM User LEFT JOIN UserRole ON User.ID = UserRole.UserID
	WHERE UserRole.ID IS NOT NULL
	`

	admins, err := ORM.FetchAll[*User](sql)
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
