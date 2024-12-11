package ORM

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Database"
	"github.com/superg3m/stoic-go/Core/Utility"
)

type InterfaceCanCRUD interface {
	canCreate() bool
	canUpdate() bool
	canDelete() bool
}

type StoicModel struct {
	InterfaceCanCRUD

	DB        *sqlx.DB
	TableName string
	isCreated bool
}

func (model *StoicModel) Update() {
	Utility.AssertMsg(model.DB != nil, fmt.Sprintf("%s Model have a valid DB connection for table: %s", model.TableName))
	Utility.AssertMsg(model.isCreated, fmt.Sprintf("%s Model must be created first before attempting to update!", model.TableName))
	Utility.AssertMsg(model.InterfaceCanCRUD.canUpdate(), "canUpdate() returned false")

	MemberNames := Utility.GetStructMemberNames(model)
	for _, memberName := range MemberNames {
		fieldMeta, exists := getAttribute(model.TableName, memberName)
		Utility.Assert(exists)
		Utility.AssertMsg(fieldMeta.isUpdatable(), fmt.Sprintf("field '%s' is not updatable", memberName))
	}

	err := Database.UpdateRecord(model.DB, model.TableName, model)
	Utility.AssertOnError(err)
}

func (model *StoicModel) Create() {
	Utility.AssertMsg(model.DB != nil, fmt.Sprintf("%s Model have a valid DB connection for table: %s", model.TableName))
	Utility.AssertMsg(model.InterfaceCanCRUD.canCreate(), "canCreate() returned false")

	MemberNames := Utility.GetStructMemberNames(model)
	for _, memberName := range MemberNames {
		_, exists := getAttribute(model.TableName, memberName)
		Utility.Assert(exists)
	}

	model.isCreated = true

	err := Database.InsertRecord(model.DB, model.TableName, model) // THIS MUST SET THE PRIMARY KEY!

	// figureOut primary keys and sure you update them accordingly for example if its an ID then you need to
	// do the last generated id from sql!

	Utility.AssertOnError(err)
}

func (model *StoicModel) Delete() {
	Utility.AssertMsg(model.DB != nil, fmt.Sprintf("%s Model have a valid DB connection for table: %s", model.TableName))
	Utility.AssertMsg(model.InterfaceCanCRUD.canDelete(), "canDelete() returned false")
	Utility.AssertMsg(model.isCreated, fmt.Sprintf("%s Model must be created first before attempting to delete!", model.TableName))

	err := Database.DeleteRecord(model.DB, model.TableName, model)
	Utility.AssertOnError(err)
}
