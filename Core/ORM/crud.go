package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
)

type InterfaceCRUD interface {
	CanCreate() bool
	CanUpdate() bool
	CanDelete() bool
	Create()
	Update()
	Delete()
}

func Update[T InterfaceCRUD](model *T) {
	Utility.AssertMsg((*model).CanUpdate(), "CanUpdate() returned false")

	_, err := UpdateRecord(GetInstance(), model)
	Utility.AssertOnError(err)
}

func Create[T InterfaceCRUD](model *T) {
	Utility.AssertMsg((*model).CanCreate(), "CanCreate() returned false")

	MemberNames := Utility.GetStructMemberNames(*model)
	hasAutoIncrement := false
	tableName := Utility.GetTypeName(*model)
	for _, memberName := range MemberNames {
		attribute, exists := getAttribute(tableName, memberName)

		if attribute.isAutoIncrement() {
			hasAutoIncrement = true
		}

		Utility.Assert(exists)
	}

	result, err := InsertRecord(GetInstance(), model)

	if hasAutoIncrement && err == nil {
		id, _ := result.LastInsertId()
		Utility.UpdateMemberValue(model, "ID", id)
	}

	// Ensure the primary key is updated, e.g., retrieve the last generated ID if applicable
	Utility.AssertOnError(err) // This doesn't make sense I should return an error code instead
}

func Delete[T InterfaceCRUD](model *T) {
	Utility.AssertMsg((*model).CanDelete(), "CanDelete() returned false")

	_, err := DeleteRecord(GetInstance(), model)
	Utility.AssertOnError(err)
}
