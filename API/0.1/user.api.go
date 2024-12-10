package API

import (
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/User"
	"time"
)

func createUser(request *Router.StoicRequest, response Router.StoicResponse) {
	username := request.GetStringParam("username")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user := &User.User{}
	user.Username = username
	user.Email = email
	user.Password = password
	user.Joined = time.Now()
	user.Create() // Creates in the database

	// w.SetData(data)
}

func sendUserMetrics(request *Router.StoicRequest, response Router.StoicResponse) {
	// w.SetData(data)
}

func init() {
	Router.RegisterApiEndpoint("User/Create", createUser, "POST",
		Router.MiddlewareValidParams("username", "email", "password"),
	)

	Router.RegisterApiEndpoint("User/Metric", sendUserMetrics, "POST")
}
