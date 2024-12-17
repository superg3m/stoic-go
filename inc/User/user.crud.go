package User

import (
	"reflect"

	"github.com/superg3m/stoic-go/Core/ORM"
)

func (u User) CanCreate() bool {
	return true
}

func (u User) CanUpdate() bool {
	return true
}

func (u User) CanDelete() bool {
	return true
}

func (u User) Create() {
	ORM.Create(&u)
}

func (u User) Update() {
	ORM.Update(&u)
}

func (u User) Delete() {
	ORM.Delete(&u)
}

func (u User) SetCache() {
	checkUser = u
}

func (u User) GetCacheDiff() []string {
	var mismatchedFields []string

	v1 := reflect.ValueOf(u)
	v2 := reflect.ValueOf(checkUser)
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

var _ ORM.InterfaceCRUD = User{}
var checkUser User
