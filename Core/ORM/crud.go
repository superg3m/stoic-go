package ORM

import (
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
	"github.com/superg3m/stoic-go/Core/Utility"
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

func extractModelComponents[T InterfaceCRUD](model *T) (StoicModel, InterfaceCRUD) {
	val := reflect.ValueOf(model)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		Utility.AssertMsg(false, "Model must be a non-nil pointer")
	}

	elem := val.Elem()
	stoicModelField := elem.FieldByName("StoicModel")

	stoicModel, ok := stoicModelField.Interface().(StoicModel)
	Utility.Assert(ok)

	return stoicModel, *model
}

func Update[T InterfaceCRUD](model *T) {
	stoicModel, crud := extractModelComponents(model)

	Utility.AssertMsg(stoicModel.DB != nil, fmt.Sprintf("%s Model must have a valid DB connection", stoicModel.TableName))
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

func Create[T InterfaceCRUD](model *T) {
	stoicModel, crud := extractModelComponents(model)

	Utility.AssertMsg(stoicModel.DB != nil, fmt.Sprintf("Model must have a valid DB connection for table: %s", stoicModel.TableName))
	Utility.AssertMsg(crud.CanCreate(), "canCreate() returned false")

	MemberNames := Utility.GetStructMemberNames(*model)
	for _, memberName := range MemberNames {
		attribute, exists := getAttribute(stoicModel.TableName, memberName)

		if attribute.isAutoIncrement() {
			sql := "SELECT MAX(?) FROM ?"
			primaryKey := -1 // this should be updating model.field
			GetInstance().QueryRowx(sql, attribute.ColumnName, stoicModel.TableName).Scan(&primaryKey)
			primaryKey += 1
		}

		Utility.Assert(exists)
	}

	stoicModel.isCreated = true

	err := InsertRecord(stoicModel.DB, stoicModel.TableName, model)

	// Ensure the primary key is updated, e.g., retrieve the last generated ID if applicable
	Utility.AssertOnError(err) // This doesn't make sense I should return an error code instead
	// maybe I should assert here but only if create is gaurenteed to succeed which mean that
	// can Create should cover that stuff
}

func Delete[T InterfaceCRUD](model *T) {
	stoicModel, crud := extractModelComponents(model)

	Utility.AssertMsg(stoicModel.DB != nil, fmt.Sprintf("%s Model must have a valid DB connection", stoicModel.TableName))
	Utility.AssertMsg(stoicModel.isCreated, fmt.Sprintf("%s Model must be created first before attempting to delete!", stoicModel.TableName))
	Utility.AssertMsg(crud.CanDelete(), "canDelete() returned false")

	err := DeleteRecord(stoicModel.DB, stoicModel.TableName, model)
	Utility.AssertOnError(err)
}
