package User

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
	"time"
)


type User struct {
	DB *sqlx.DB