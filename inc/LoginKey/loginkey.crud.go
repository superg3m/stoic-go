package LoginKey

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"reflect"
)

func (model *LoginKey) CanCreate() bool {
	return true
}

func (model *LoginKey) CanRead() bool {
	return true
}

func (model *LoginKey) CanUpdate() bool {
	return true
}

func (model *LoginKey) CanDelete() bool {
	return true
}

func (model *LoginKey) Create() ORM.CrudReturn {
	return ORM.Create(&model)
}

func (model *LoginKey) Read() ORM.CrudReturn {
	return ORM.Read(&model)
}

func (model *LoginKey) Update() ORM.CrudReturn {
	return ORM.Update(&model)
}

func (model *LoginKey) Delete() ORM.CrudReturn {
	return ORM.Delete(&model)
}

func (u *LoginKey) SetCache() {
	cache = *u
}

func (u LoginKey) GetCacheDiff() []string {
	var mismatchedFields []string

	v1 := reflect.ValueOf(u)
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
