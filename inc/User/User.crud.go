package User

import (
	"github.com/superg3m/stoic-go/Core/Utility"
	"reflect"

	"github.com/superg3m/stoic-go/Core/ORM"
)

func (u *User) CanCreate() []string {
	var errors []string = nil
	if !Utility.ValidEmail(u.Email) {
		errors = append(errors, "User Invalid Email")
	}

	dbUser := *u
	read := dbUser.Read()
	if read.IsBad() {
		return errors
	}

	if u.Email == dbUser.Email {
		errors = append(errors, "User Duplicate Email")
	}

	return errors
}

func (u *User) CanRead() []string {
	if u.ID > 0 {
		return nil
	}

	return nil
}

func (u *User) CanUpdate() []string {
	dbUser := *u
	read := dbUser.Read()
	if read.IsBad() {
		return nil
	}

	var errors []string = nil
	if !Utility.ValidEmail(u.Email) {
		errors = append(errors, "User Invalid Email")
	}

	if u.ID == dbUser.ID {
		return nil
	}

	if u.Email == dbUser.Email {
		errors = append(errors, "User Duplicate Email")
	}

	return errors
}

func (u *User) CanDelete() []string {
	return nil
}

func (u *User) Create() ORM.CrudReturn {
	return ORM.Create(u, u.DB)
}

func (u *User) Read() ORM.CrudReturn {
	return ORM.Read(u, u.DB)
}

func (u *User) Update() ORM.CrudReturn {
	return ORM.Update(u, u.DB)
}

func (u *User) Delete() ORM.CrudReturn {
	return ORM.Delete(u, u.DB)
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
