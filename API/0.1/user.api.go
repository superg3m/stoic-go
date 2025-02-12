package API

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/Utility"
	"time"

	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/LoginKey"
	"github.com/superg3m/stoic-go/inc/User"
)

func createUser(request *Router.StoicRequest, response Router.StoicResponse) {
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user := User.New()
	user.Email = email
	create := user.Create()
	if create.IsBad() {
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

	response.SetData(user)
}

func updateUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")
	emailConfirmed := request.GetBoolParam("emailConfirmed")

	user, err := User.FromID(id)
	if err != nil {
		response.AddError(err.Error())
		return
	}

	user.Email = email
	user.EmailConfirmed = emailConfirmed
	user.LastLogin = Utility.NewTime(time.Now())
	update := user.Update()
	if update.IsBad() {
		response.SetError("Failed to update user | %s", update.GetError())
		return
	}

	loginKey, err := LoginKey.FromUserID_Provider(user.ID, LoginKey.PASSWORD)
	if err != nil {
		response.SetError(err.Error())
		return
	}
	loginKey.UserID = user.ID
	loginKey.Key = password
	loginKey.Provider = LoginKey.PASSWORD
	loginKey.HashKey()
	update = loginKey.Update()
	if update.IsBad() {
		response.SetError("Failed to update login key | %s", update.GetError())
		return
	}

	response.SetData("User Updated Successfully!")
}

func deleteUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")

	user, err := User.FromID(id)
	if err != nil {
		response.SetError(err.Error())
		return
	}

	del := user.Delete()
	if del.IsBad() {
		response.SetError("Failed to delete user | %s", del.GetError())
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
