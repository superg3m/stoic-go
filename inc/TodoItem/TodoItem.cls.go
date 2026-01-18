package TodoItem

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

type TodoItem struct {
	DB *sqlx.DB `json:"-"`

	ID      int    `json:"ID"`
	OwnerID int    `json:"OwnerID"`
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

var DatabaseName = "stoic"

func New() *TodoItem {
	ret := new(TodoItem)

	ret.DB = ORM.GetInstance(DatabaseName)
	ret.ID = 0
	ret.OwnerID = 0
	ret.Message = ""
	ret.Status = -1

	return ret
}

func FromID(ID int) (*TodoItem, []string) {
	ret := New()
	ret.ID = ID
	read := ret.Read()
	if read.IsBad() {
		return nil, read.GetErrors()
	}

	return ret, nil
}

func AllFromOwnerID(OwnerID int) ([]*TodoItem, error) {
	sql := "SELECT * From TodoItem WHERE OwnerID = ?"
	todos, err := ORM.FetchAll[*TodoItem](ORM.GetInstance(DatabaseName), sql, OwnerID)
	return todos, err
}

func init() {
	ORM.RegisterTableName(&TodoItem{})
	ORM.RegisterTableColumn("ID", "ID", ORM.AUTO_INCREMENT|ORM.KEY)
	ORM.RegisterTableColumn("OwnerID", "OwnerID")
	ORM.RegisterTableColumn("Message", "Message", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Status", "Status", ORM.UPDATABLE)
}
