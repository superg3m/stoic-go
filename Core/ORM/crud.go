package ORM

import (
	"github.com/jmoiron/sqlx"
	"slices"

	"github.com/superg3m/stoic-go/Core/Utility"
)

type InterfaceCRUD interface {
	CanCreate() []string
	CanRead() []string
	CanUpdate() []string
	CanDelete() []string

	Create() CrudReturn
	Read() CrudReturn
	Update() CrudReturn
	Delete() CrudReturn

	SetCache()
	GetCacheDiff() []string
}

var excludeList = []string{"DB"}

func Create[T InterfaceCRUD](model T, db *sqlx.DB) CrudReturn {
	ret := CreateCRUD()

	errors := model.CanCreate()
	if errors != nil {
		ret.AddErrors(errors)
		return ret
	}

	payload := getModelPayload(model)
	hasAutoIncrement := createValidate(payload)

	result, err := CreateRecord(db, payload)
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

func Read[T InterfaceCRUD](model T, db *sqlx.DB) CrudReturn {
	ret := CreateCRUD()

	errors := model.CanRead()
	if errors != nil {
		ret.AddErrors(errors)
		return ret
	}

	payload := getModelPayload(model)

	err := ReadRecord(db, payload, model)
	if err != nil {
		ret.AddError(err.Error())
		return ret
	}

	model.SetCache()

	return ret
}

func Update[T InterfaceCRUD](model T, db *sqlx.DB) CrudReturn {
	ret := CreateCRUD()

	errors := model.CanUpdate()
	if errors != nil {
		ret.AddErrors(errors)
		return ret
	}

	payload := getModelPayload(model)
	updateValidate(payload, model)

	_, err := UpdateRecord(db, payload)
	if err != nil {
		ret.AddError(err.Error())
		return ret
	}

	model.SetCache()

	return ret
}

func Delete[T InterfaceCRUD](model T, db *sqlx.DB) CrudReturn {
	ret := CreateCRUD()

	errors := model.CanDelete()
	if errors != nil {
		ret.AddErrors(errors)
		return ret
	}

	read := Read(model, db)
	if read.IsBad() {
		ret.AddErrors(ret.GetErrors(), "Failed to delete")
		return ret
	}

	payload := getModelPayload(model)
	_, err := DeleteRecord(db, payload)
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
