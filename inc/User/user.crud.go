package User

import "github.com/superg3m/stoic-go/Core/ORM"

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

var _ ORM.InterfaceCRUD = User{}
