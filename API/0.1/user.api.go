package API

import (
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/Core/Utility"
	"github.com/superg3m/stoic-go/inc/User"
	"time"
)

func createUser(request *Router.StoicRequest, response Router.StoicResponse) {
	username := request.GetStringParam("username")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user := User.New()
	user.Username = username
	user.Password = password
	user.Email = email
	create := user.Create()
	if create.IsBad() {
		response.SetError("Failed to create user | %s", create.GetError())
		return
	}

	response.SetData("User Created Successfully!")
}

func updateUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")
	username := request.GetStringParam("username")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")
	emailConfirmed := request.GetBoolParam("emailConfirmed")

	user, err := User.FromID(id)
	if err != nil {
		response.SetError(err.Error())
		return
	}

	user.Username = username
	user.Email = email
	user.Password = password
	user.EmailConfirmed = emailConfirmed
	user.LastLogin = Utility.NewTime(time.Now())
	update := user.Update()
	if update.IsBad() {
		response.SetError("Failed to update user | %s", update.GetError())
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

func sendUserMetrics(request *Router.StoicRequest, response Router.StoicResponse) {
	// w.SetData(data)
}

func init() {
	Router.RegisterApiEndpoint("/User/Create", createUser, "POST",
		Router.MiddlewareValidParams("username", "email", "password"),
	)

	Router.RegisterApiEndpoint("/User/Update", updateUser, "POST",
		Router.MiddlewareValidParams("id", "username", "email", "password", "emailConfirmed"),
	)

	Router.RegisterApiEndpoint("/User/Delete", deleteUser, "POST",
		Router.MiddlewareValidParams("id"),
	)

	Router.RegisterApiEndpoint("/User/Metric", sendUserMetrics, "POST")
}
