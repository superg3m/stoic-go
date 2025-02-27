package User

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"reflect"
)

func (model *User) Create() ORM.CrudReturn {
	return ORM.Create(model)
}

func (model *User) Read() ORM.CrudReturn {
	return ORM.Read(model)
}

func (model *User) Update() ORM.CrudReturn {
	return ORM.Update(model)
}

func (model *User) Delete() ORM.CrudReturn {
	return ORM.Delete(model)
}

func (model *User) SetCache() {
	cache = *model
}

func (model *User) GetCacheDiff() []string {
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

var _ ORM.InterfaceCRUD = &User{}
var cache User