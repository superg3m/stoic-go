package ORM

import (
	"slices"

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

func Create[T InterfaceCRUD](model T) CrudReturn {
	ret := CreateCRUD()
	if !model.CanCreate() {
		ret.AddError("CanCreate() returned false")
		return ret
	}

	payload := getModelPayload(model)
	hasAutoIncrement := createValidate(payload)

	result, err := CreateRecord(GetInstance(), payload)
	if err != nil {
		ret.AddError(err.Error())
		return ret
	}

	if hasAutoIncrement {
		id, _ := result.LastInsertId()
		Utility.UpdateMemberValue(model, "ID", id)
	}

	model.SetCache()

	return ret
}

func Read[T InterfaceCRUD](model T) CrudReturn {
	ret := CreateCRUD()
	if !model.CanRead() {
		ret.AddError("CanRead() returned false")
		return ret
	}

	payload := getModelPayload(model)

	err := ReadRecord(GetInstance(), payload, model)
	if err != nil {
		ret.AddError(err.Error())
		return ret
	}

	model.SetCache()

	return ret
}

func Update[T InterfaceCRUD](model T) CrudReturn {
	ret := CreateCRUD()
	if !model.CanUpdate() {
		ret.AddError("CanUpdate() returned false")
		return ret
	}

	payload := getModelPayload(model)
	updateValidate(payload, model)

	_, err := UpdateRecord(GetInstance(), payload)
	if err != nil {
		ret.AddError(err.Error())
		return ret
	}

	model.SetCache()

	return ret
}

func Delete[T InterfaceCRUD](model T) CrudReturn {
	ret := CreateCRUD()
	Utility.AssertMsg(model.CanDelete(), "CanDelete() returned false")

	read := Read(model)
	if read.IsBad() {
		ret.AddErrors(ret.GetErrors(), "Failed to delete")
		return ret
	}

	payload := getModelPayload(model)
	_, err := DeleteRecord(GetInstance(), payload)
	if err != nil {
		ret.AddError(err.Error())
		return ret
	}

	model.SetCache()

	return ret
}

// returns if auto increment
func createValidate(payload ModelPayload) bool {
	hasAutoIncrement := false

	for _, attribute := range GetAttributes(payload.TableName) {
		if attribute.isAutoIncrement() {
			hasAutoIncrement = true
		}
	}

	return hasAutoIncrement
}

func updateValidate[T InterfaceCRUD](payload ModelPayload, model T) {
	membersChanged := (model).GetCacheDiff()

	for memberName, attribute := range GetAttributes(payload.TableName) {
		if !slices.Contains(membersChanged, memberName) {
			continue
		}
		Utility.AssertMsg(attribute.isUpdatable(), "%s.%s is not updatable", payload.TableName, memberName)
	}
}
