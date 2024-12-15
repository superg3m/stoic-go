package API

import (
	"github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/User"
)

func createUser(request *Router.StoicRequest, response Router.StoicResponse) {
	username := request.GetStringParam("username")
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user := User.New()
	user.ID = 2
	user.Username = username
	user.Email = email
	user.Password = password
	// user.Joined = time.Now()
	ORM.Create(user)

	user.Delete()

	response.SetData("User Created Successfully!")
}

func sendUserMetrics(request *Router.StoicRequest, response Router.StoicResponse) {
	// w.SetData(data)
}

func init() {
	Router.RegisterApiEndpoint("/User/Create", createUser, "POST",
		Router.MiddlewareValidParams("username", "email", "password"),
	)

	Router.RegisterApiEndpoint("/User/Metric", sendUserMetrics, "POST")
}
