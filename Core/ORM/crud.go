package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
)

type InterfaceCRUD interface {
	CanCreate() bool
	CanRead() bool
	CanUpdate() bool
	CanDelete() bool

	Create() bool
	Read() bool
	Update() bool
	Delete() bool

	SetCache()
	GetCacheDiff() []string
}

var excludeList = []string{"DB"}

func Create[T InterfaceCRUD](model *T) bool {
	Utility.AssertMsg((*model).CanCreate(), "CanCreate() returned false")

	MemberNames := Utility.GetStructMemberNames(*model, excludeList...)
	hasAutoIncrement := false
	tableName := Utility.GetTypeName(*model)
	for _, memberName := range MemberNames {
		attribute, exists := getAttribute(tableName, memberName)

		if attribute.isAutoIncrement() {
			hasAutoIncrement = true
		}

		Utility.Assert(exists)
	}

	result, err := CreateRecord(GetInstance(), model)

	if hasAutoIncrement && err == nil {
		id, _ := result.LastInsertId()
		Utility.UpdateMemberValue(model, "ID", id)
	}

	(*model).SetCache()

	return err == nil
}

func Read[T InterfaceCRUD](model *T) bool {
	Utility.AssertMsg((*model).CanRead(), "CanRead() returned false")

	err := ReadRecord(GetInstance(), model)
	return err == nil
}

func Update[T InterfaceCRUD](model *T) bool {
	Utility.AssertMsg((*model).CanUpdate(), "CanUpdate() returned false")

	tableName := Utility.GetTypeName(*model)
	membersChanged := (*model).GetCacheDiff()

	for _, member := range membersChanged {
		attribute, _ := getAttribute(tableName, member)
		Utility.AssertMsg(attribute.isUpdatable(), "%s.%s is not updatable", tableName, member)
	}

	_, err := UpdateRecord(GetInstance(), model)
	return err == nil
}

func Delete[T InterfaceCRUD](model *T) bool {
	Utility.AssertMsg((*model).CanDelete(), "CanDelete() returned false")

	_, err := DeleteRecord(GetInstance(), model)
	return err == nil
}
