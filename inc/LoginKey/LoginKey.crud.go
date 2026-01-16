package LoginKey

import (
	"reflect"

	"github.com/superg3m/stoic-go/Core/ORM"
)

func (model *LoginKey) CanCreate() []string {
	return nil
}

func (model *LoginKey) CanRead() []string {
	return nil
}

func (model *LoginKey) CanUpdate() []string {
	return nil
}

func (model *LoginKey) CanDelete() []string {
	return nil
}

func (model *LoginKey) Create() ORM.CrudReturn {
	return ORM.Create(model, model.DB)
}

func (model *LoginKey) Read() ORM.CrudReturn {
	return ORM.Read(model, model.DB)
}

func (model *LoginKey) Update() ORM.CrudReturn {
	return ORM.Update(model, model.DB)
}

func (model *LoginKey) Delete() ORM.CrudReturn {
	return ORM.Delete(model, model.DB)
}

func (model *LoginKey) SetCache() {
	cache = *model
}

func (model *LoginKey) GetCacheDiff() []string {
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

var _ ORM.InterfaceCRUD = &LoginKey{}
var cache LoginKey
