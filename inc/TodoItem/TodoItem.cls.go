package TodoItem

import (
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/ORM"
)

type TodoItem struct {
	DB      *sqlx.DB
	ID      int
	OwnerID int
	Message string
	Status  int
}

var DatabaseName = "stoic"

func GetDBInstance() *sqlx.DB {
	return ORM.GetInstance(DatabaseName)
}

func New() *TodoItem {
	ret := new(TodoItem)

	//ret.DB = ORM.GetInstance()
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

func AllFromOwnerID(OwnerID int) []*TodoItem {
	sql := "SELECT * From TodoItem WHERE OwnerID = ?"
	todos, _ := ORM.FetchAll[*TodoItem](GetDBInstance(), sql, OwnerID)
	return todos
}

func init() {
	ORM.RegisterTableName(&TodoItem{})
	ORM.RegisterTableColumn("ID", "ID", ORM.AUTO_INCREMENT|ORM.KEY)
	ORM.RegisterTableColumn("OwnerID", "OwnerID")
	ORM.RegisterTableColumn("Message", "Message", ORM.UPDATABLE)
	ORM.RegisterTableColumn("Status", "Status", ORM.UPDATABLE)
}
