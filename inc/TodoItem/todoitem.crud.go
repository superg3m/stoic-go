package TodoItem

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"reflect"
)

func (model *TodoItem) CanCreate() []string {
	return nil
}

func (model *TodoItem) CanRead() []string {
	return nil
}

func (model *TodoItem) CanUpdate() []string {
	return nil
}

func (model *TodoItem) CanDelete() []string {
	return nil
}

func (model *TodoItem) Create() ORM.CrudReturn {
	return ORM.Create(model)
}

func (model *TodoItem) Read() ORM.CrudReturn {
	return ORM.Read(model)
}

func (model *TodoItem) Update() ORM.CrudReturn {
	return ORM.Update(model)
}

func (model *TodoItem) Delete() ORM.CrudReturn {
	return ORM.Delete(model)
}

func (model *TodoItem) SetCache() {
	cache = *model
}

func (model *TodoItem) GetCacheDiff() []string {
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

var _ ORM.InterfaceCRUD = &TodoItem{}
var cache TodoItem