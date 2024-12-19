package API

import (
    "fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/User"
)

func GetUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")

	record, err := User.FromId(id)
	if err != nil {
		response.SetError(fmt.Sprintf("Error fetching User: %s", err))
		return
	}

	response.SetData(record)
}

func CreateUser(request *Router.StoicRequest, response Router.StoicResponse) {
	// Instantiate a new model object
	entity := User.New()
    entity.ID = request.GetIntParam("ID")
    entity.Username = request.GetStringParam("Username")
    request.GetJsonParam("Joined", &entity.Joined)
    entity.Email = request.GetStringParam("Email")

	entity.Create()
	response.SetData(fmt.Sprintf("User created successfully with ID %d", entity.ID))
}

func UpdateUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")

	entity, err := User.FromId(id)
	if err != nil {
		response.SetError(fmt.Sprintf("Error fetching User: %s", err))
		return
	}
    entity.ID = request.GetIntParam("ID")
    entity.Username = request.GetStringParam("Username")
    request.GetJsonParam("Joined", &entity.Joined)
    entity.Email = request.GetStringParam("Email")

	entity.Update()

	response.SetData(fmt.Sprintf("User with ID %d updated successfully", id))
}

func DeleteUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")

	err := User.DeleteByID(id) // Dynamic DeleteByID
	response.SetData(fmt.Sprintf("User with ID %d deleted successfully", id))
}

func init() {
	Router.RegisterApiEndpoint("User/Get", GetUser, "GET",
		Router.MiddlewareValidParams("id"),
	)
	Router.RegisterApiEndpoint("User/Create", CreateUser, "POST",
		Router.MiddlewareValidParams("ID", "Username", "Joined", "Email", ),
	)
	Router.RegisterApiEndpoint("User/Update", UpdateUser, "PATCH",
		Router.MiddlewareValidParams("id", "ID", "Username", "Joined", "Email", ),
	)
	Router.RegisterApiEndpoint("User/Delete", DeleteUser, "DELETE",
		Router.MiddlewareValidParams("id"),
	)
}