package ORM

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"
)

type InterfaceCRUD interface {
	CanCreate() bool
	CanUpdate() bool
	CanDelete() bool
}

type StoicModel struct {
	DB        *sqlx.DB
	TableName string
	isCreated bool
}

func extractModelComponents(model any) (StoicModel, InterfaceCRUD) {
	val := reflect.ValueOf(model)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		Utility.AssertMsg(false, "Model must be a non-nil pointer")
	}

	elem := val.Elem()
	stoicModelField := elem.FieldByName("StoicModel")

	stoicModel, ok := stoicModelField.Interface().(StoicModel)
	Utility.Assert(ok)

	// Check if the model implements the InterfaceCRUD interface
	crud, ok := model.(InterfaceCRUD)
	Utility.Assert(ok)

	return stoicModel, crud
}

func Update(model any) {
	stoicModel, crud := extractModelComponents(model)
	model = Utility.DereferencePointer(model)

	Utility.AssertMsg(stoicModel.DB != nil, fmt.Sprintf("%s Model must have a valid DB connection for table: %s", stoicModel.TableName))
	Utility.AssertMsg(stoicModel.isCreated, fmt.Sprintf("%s Model must be created first before attempting to update!", stoicModel.TableName))
	Utility.AssertMsg(crud.CanUpdate(), "canUpdate() returned false")

	MemberNames := Utility.GetStructMemberNames(model)
	for _, memberName := range MemberNames {
		fieldMeta, exists := getAttribute(stoicModel.TableName, memberName)
		Utility.Assert(exists)
		Utility.AssertMsg(fieldMeta.isUpdatable(), fmt.Sprintf("field '%s' is not updatable", memberName))
	}

	err := UpdateRecord(stoicModel.DB, stoicModel.TableName, model)
	Utility.AssertOnError(err)
}

func Create(model any) {
	stoicModel, crud := extractModelComponents(model)
	model = Utility.DereferencePointer(model)

	Utility.AssertMsg(stoicModel.DB != nil, fmt.Sprintf("Model must have a valid DB connection for table: %s", stoicModel.TableName))
	Utility.AssertMsg(crud.CanCreate(), "canCreate() returned false")

	MemberNames := Utility.GetStructMemberNames(model)
	for _, memberName := range MemberNames {
		_, exists := getAttribute(stoicModel.TableName, memberName)
		Utility.Assert(exists)
	}

	stoicModel.isCreated = true

	err := InsertRecord(stoicModel.DB, stoicModel.TableName, model) // THIS MUST SET THE PRIMARY KEY!

	// Ensure the primary key is updated, e.g., retrieve the last generated ID if applicable
	Utility.AssertOnError(err)
}

func Delete(model any) {
	stoicModel, crud := extractModelComponents(model)
	model = Utility.DereferencePointer(model)

	Utility.AssertMsg(stoicModel.DB != nil, fmt.Sprintf("%s Model must have a valid DB connection for table: %s", stoicModel.TableName))
	Utility.AssertMsg(stoicModel.isCreated, fmt.Sprintf("%s Model must be created first before attempting to delete!", stoicModel.TableName))
	Utility.AssertMsg(crud.CanDelete(), "canDelete() returned false")

	err := DeleteRecord(stoicModel.DB, stoicModel.TableName, model)
	Utility.AssertOnError(err)
}
