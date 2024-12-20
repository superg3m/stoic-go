package User

import (
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"

	"github.com/superg3m/stoic-go/Core/ORM"
)

func (u *User) CanCreate() bool {
	if !Utility.ValidEmail(u.Email) {
		return false
	}

	return true
}

func (u *User) CanRead() bool {
	if u.ID >= 1 {
		return true
	}

	if Utility.ValidEmail(u.Email) {
		return true
	}

	return false
}

func (u *User) CanUpdate() bool {
	return true
}

func (u *User) CanDelete() bool {
	return true
}

func (u *User) Create() ORM.CrudReturn {
	return ORM.Create(u)
}

func (u *User) Read() ORM.CrudReturn {
	return ORM.Read(u)
}

func (u *User) Update() ORM.CrudReturn {
	return ORM.Update(u)
}

func (u *User) Delete() ORM.CrudReturn {
	return ORM.Delete(u)
}

func (u *User) SetCache() {
	cache = *u
}

func (u *User) GetCacheDiff() []string {
	var mismatchedFields []string

	v1 := reflect.ValueOf(*u)
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
