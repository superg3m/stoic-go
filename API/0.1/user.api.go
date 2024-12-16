package API

import (
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/User"
)

func createUser(request *Router.StoicRequest, response Router.StoicResponse) {
	username := request.GetStringParam("username")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user := User.New()
	user.Username = username
	user.Email = email
	user.Password = password
	// user.Joined = time.Now()
	user.Create()

	//ORM.Delete(user)

	response.SetData("User Created Successfully!")
}

func updateUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")
	username := request.GetStringParam("username")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")
	emailConfirmed := request.GetBoolParam("emailConfirmed")

	user := User.FromID(id)
	user.Username = username
	user.Email = email
	user.Password = password
	user.EmailConfirmed = emailConfirmed
	user.Update()

	response.SetData("User Updated Successfully!")
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

	Router.RegisterApiEndpoint("/User/Metric", sendUserMetrics, "POST")
}
