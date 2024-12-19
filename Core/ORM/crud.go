package ORM

import (
	"github.com/superg3m/stoic-go/Core/Utility"
)

type InterfaceCRUD interface {
	CanCreate() bool
	CanRead() bool
	CanUpdate() bool
	CanDelete() bool

	Create() CrudReturn
	Read() CrudReturn
	Update() CrudReturn
	Delete() CrudReturn

	SetCache()
	GetCacheDiff() []string
}

var excludeList = []string{"DB"}

func Create[T InterfaceCRUD](model *T) CrudReturn {
	ret := CreateCRUD()
	if !(*model).CanCreate() {
		ret.setErrorMsg("CanCreate() returned false")
		return ret
	}

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
	if err != nil {
		ret.setErrorMsg(err.Error())
		return ret
	}

	if hasAutoIncrement && err == nil {
		id, _ := result.LastInsertId()
		Utility.UpdateMemberValue(model, "ID", id)
	}

	(*model).SetCache()

	return ret
}

func Read[T InterfaceCRUD](model *T) CrudReturn {
	ret := CreateCRUD()
	if !(*model).CanRead() {
		ret.setErrorMsg("CanRead() returned false")
		return ret
	}

	err := ReadRecord(GetInstance(), model)
	if err != nil {
		ret.setErrorMsg(err.Error())
		return ret
	}

	return ret
}

func Update[T InterfaceCRUD](model *T) CrudReturn {
	ret := CreateCRUD()
	if !(*model).CanUpdate() {
		ret.setErrorMsg("CanUpdate() returned false")
		return ret
	}

	tableName := Utility.GetTypeName(*model)
	membersChanged := (*model).GetCacheDiff()

	for _, member := range membersChanged {
		attribute, _ := getAttribute(tableName, member)
		Utility.AssertMsg(attribute.isUpdatable(), "%s.%s is not updatable", tableName, member)
	}

	_, err := UpdateRecord(GetInstance(), model)
	if err != nil {
		ret.setErrorMsg(err.Error())
		return ret
	}

	return ret
}

func Delete[T InterfaceCRUD](model *T) CrudReturn {
	ret := CreateCRUD()
	Utility.AssertMsg((*model).CanDelete(), "CanDelete() returned false")

	_, err := DeleteRecord(GetInstance(), model)
	if err != nil {
		ret.setErrorMsg(err.Error())
		return ret
	}

	return ret
}
