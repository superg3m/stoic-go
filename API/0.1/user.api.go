package API

import (
    "fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/User"
)

func createUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	entity := User.New()
    entity.ID = request.GetIntParam("ID")


    create := entity.Create()
    if create.IsBad() {
        response.AddErrors(create.GetErrors(), "Failed to create User")
        return
    }

	response.SetData(fmt.Sprintf("User created successfully"))
}

func getUser(request *Router.StoicRequest, response *Router.StoicResponse) {
    ID := request.GetIntParam("ID")

	entity, errors := User.FromID(ID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get User")
		return
	}

	response.SetData(entity)
}

func updateUser(request *Router.StoicRequest, response *Router.StoicResponse) {
    ID := request.GetIntParam("ID")

	entity, errors := User.FromID(ID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get User")
		return
	}
    entity.ID = request.GetIntParam("ID")
    entity.Email = request.GetStringParam("Email")
    entity.EmailConfirmed = request.GetBoolParam("EmailConfirmed")

	update := entity.Update()
	if update.IsBad() {
    	response.AddErrors(update.GetErrors(), "Failed to update User")
	    return
	}

	response.SetData(fmt.Sprintf("User updated successfully"))
}

func deleteUser(request *Router.StoicRequest, response *Router.StoicResponse) {
    ID := request.GetIntParam("ID")

	entity, errors := User.FromID(ID)
	if errors != nil {
	    response.AddErrors(errors, "Failed to get User")
		return
	}

	del := entity.Delete()

	if del.IsBad() {
	    response.AddErrors(del.GetErrors(), "Failed to delete User")
	    return
	}

	response.SetData(fmt.Sprintf("User deleted successfully"))
}

func init() {
	Router.RegisterApiEndpoint("User/Create", createUser, "POST",
		Router.MiddlewareValidParams("ID", "Email", "EmailConfirmed", "Joined", "LastLogin", "LastActive"),
	)
    Router.RegisterApiEndpoint("User/Get", getUser, "GET",
        Router.MiddlewareValidParams("ID"),
    )
	Router.RegisterApiEndpoint("User/Update", updateUser, "PATCH",
		Router.MiddlewareValidParams("ID", "Email", "EmailConfirmed", "Joined", "LastLogin", "LastActive"),
	)
	Router.RegisterApiEndpoint("User/Delete", deleteUser, "DELETE",
		Router.MiddlewareValidParams("ID"),
	)
}