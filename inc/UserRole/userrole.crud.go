package UserRole

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"reflect"
)

func (model *UserRole) CanCreate() []string {
	model, errors := FromUserID_RoleID(model.UserID, model.RoleID)
	if errors != nil {
		errors = append(errors, "Duplicate User Role")
		return errors
	}

	return nil
}

func (model *UserRole) CanRead() []string {
	return nil
}

func (model *UserRole) CanUpdate() []string {
	return nil
}

func (model *UserRole) CanDelete() []string {
	return nil
}

func (model *UserRole) Create() ORM.CrudReturn {
	return ORM.Create(model)
}

func (model *UserRole) Read() ORM.CrudReturn {
	return ORM.Read(model)
}

func (model *UserRole) Update() ORM.CrudReturn {
	return ORM.Update(model)
}

func (model *UserRole) Delete() ORM.CrudReturn {
	return ORM.Delete(model)
}

func (model *UserRole) SetCache() {
	cache = *model
}

func (model *UserRole) GetCacheDiff() []string {
	var mismatchedFields []string

	v1 := reflect.ValueOf(*model)
	v2 := reflect.ValueOf(cache)
	userType := v1.Type()

	for i := 0; i < v1.NumField(); i++ {
		f1 := v1.Field(i)
		f2 := v2.Field(i)

		if !reflect.DeepEqual(f1.Interface(), f2.Interface()) {
			mismatchedFields = append(mismatchedFields, userType.Field(i).Name)
		}
	}

	return mismatchedFields
}

var _ ORM.InterfaceCRUD = &UserRole{}
var cache UserRole
