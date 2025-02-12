package API

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/Utility"
	"time"

	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/LoginKey"
	"github.com/superg3m/stoic-go/inc/User"
)

func createUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user := User.New()
	user.Email = email
	create := user.Create()
	if create.IsBad() {
		fmt.Println("Create Endpoint")
		response.AddErrors(create.GetErrors(), "Failed to create user")

		return
	}

	loginKey := LoginKey.New()
	loginKey.UserID = user.ID
	loginKey.Key = password
	loginKey.Provider = LoginKey.PASSWORD
	loginKey.HashKey()
	create = loginKey.Create()
	if create.IsBad() {
		response.AddErrors(create.GetErrors(), "Failed to create login key")
		user.Delete()

		return
	}

	// response.SetData(user) // This should work but right now it doesn't
	response.SetData("User Created Successfully!")
}

func updateUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	id := request.GetIntParam("id")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")
	emailConfirmed := request.GetBoolParam("emailConfirmed")

	user, errors := User.FromID(id)
	if errors != nil {
		response.AddErrors(errors, "Failed to get user from ID")
		return
	}

	user.Email = email
	user.EmailConfirmed = emailConfirmed
	user.LastLogin = Utility.NewTime(time.Now())
	update := user.Update()
	if update.IsBad() {
		response.AddErrors(update.GetErrors(), "Failed to update user")
		return
	}

	loginKey, errors := LoginKey.FromUserID_Provider(user.ID, LoginKey.PASSWORD)
	if errors != nil {
		response.AddErrors(errors, "Failed to get LoginKey from UserID and Provider")
		return
	}
	loginKey.UserID = user.ID
	loginKey.Key = password
	loginKey.Provider = LoginKey.PASSWORD
	loginKey.HashKey()
	update = loginKey.Update()
	if update.IsBad() {
		response.AddErrors(update.GetErrors(), "Failed to update login key")
		return
	}

	response.SetData("User Updated Successfully!")
}

func deleteUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	id := request.GetIntParam("id")

	user, errors := User.FromID(id)
	if errors != nil {
		response.AddErrors(errors, "Failed to get user from ID")
		return
	}

	del := user.Delete()
	if del.IsBad() {
		response.AddErrors(del.GetErrors(), "Failed to delete user")
		return
	}

	response.SetData("User Deleted Successfully!")
}

func init() {
	Router.RegisterApiEndpoint("/User/Create", createUser, "POST",
		Router.MiddlewareValidParams("username", "email", "password"),
	)

	Router.RegisterApiEndpoint("/User/Update", updateUser, "POST",
		Router.MiddlewareValidParams("id", "email", "password", "emailConfirmed"),
	)

	Router.RegisterApiEndpoint("/User/Delete", deleteUser, "POST",
		Router.MiddlewareValidParams("id"),
	)
}
