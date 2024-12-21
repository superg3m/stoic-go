package ORM

import (
	"errors"
	"fmt"
	"github.com/superg3m/stoic-go/Core/Utility"
	"slices"
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
		ret.setError(errors.New("CanCreate() returned false"))
		return ret
	}

	payload := getModelPayload(model)
	hasAutoIncrement := createValidate(payload)

	result, err := CreateRecord(GetInstance(), payload)
	if err != nil {
		ret.setError(err)
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
		ret.setError(errors.New("CanRead() returned false"))
		return ret
	}

	payload := getModelPayload(model)

	err := ReadRecord(GetInstance(), payload, model)
	if err != nil {
		ret.setError(err)
		return ret
	}

	model.SetCache()

	return ret
}

func Update[T InterfaceCRUD](model T) CrudReturn {
	ret := CreateCRUD()
	if !model.CanUpdate() {
		ret.setError(errors.New("CanUpdate() returned false"))
		return ret
	}

	payload := getModelPayload(model)
	updateValidate(payload, model)

	_, err := UpdateRecord(GetInstance(), payload)
	if err != nil {
		ret.setError(err)
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
		msg := fmt.Sprintf("Failed to delete | %s", ret.GetError())
		ret.setError(errors.New(msg))
		return ret
	}

	payload := getModelPayload(model)
	_, err := DeleteRecord(GetInstance(), payload)
	if err != nil {
		ret.setError(err)
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
