package API

import (
    "fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/User"
)

func createUser(request *Router.StoicRequest, response Router.StoicResponse) {
	// Instantiate a new model object
	entity := User.New()  
        entity.ID = request.GetIntParam("ID")  
        request.GetJsonParam("Username", &entity.Username)  
        request.GetJsonParam("Email", &entity.Email)  
        request.GetJsonParam("Password", &entity.Password)  
        entity.EmailConfirmed = request.GetBoolParam("EmailConfirmed")  
        request.GetJsonParam("Joined", &entity.Joined)  
        request.GetJsonParam("LastLogin", &entity.LastLogin)  
        request.GetJsonParam("LastActive", &entity.LastActive)


    create := entity.Create()
    if create.IsBad() {
        response.SetError("Failed to create User | %s", create.GetErrorMsg())
        return
    }

	response.SetData(fmt.Sprintf("User created successfully with ID %d", entity.ID))
}

func getUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")

	entity, err := User.FromID(id)
	if err != nil {
		response.SetError(fmt.Sprintf("Error getting User: %s", err))
		return
	}

	response.SetData(entity)
}

func updateUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")

	entity, err := User.FromID(id)
    if err != nil {
        response.SetError("Failed to find User by ID | %s", err)
        return
    }
    entity.ID = request.GetIntParam("ID")
    request.GetJsonParam("Username", &entity.Username)
    request.GetJsonParam("Email", &entity.Email)
    request.GetJsonParam("Password", &entity.Password)
    entity.EmailConfirmed = request.GetBoolParam("EmailConfirmed")
    request.GetJsonParam("Joined", &entity.Joined)
    request.GetJsonParam("LastLogin", &entity.LastLogin)
    request.GetJsonParam("LastActive", &entity.LastActive)

	update := entity.Update()
	if update.IsBad() {
	    response.SetError("Failed to update User | %s", update.GetErrorMsg())
	    return
	}

	response.SetData(fmt.Sprintf("User with ID %d updated successfully", id))
}

func deleteUser(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")

	entity, err := User.FromID(id)
	if err != nil {
        response.SetError("Failed to find User by ID | %s", err)
        return
	}

	delete := entity.Delete()

	if delete.IsBad() {
	    response.SetError("Failed to delete User %s", delete.GetErrorMsg())
	    return
	}

	response.SetData(fmt.Sprintf("User with ID %d deleted successfully", id))
}

func init() {
	Router.RegisterApiEndpoint("User/Create", createUser, "POST",
		Router.MiddlewareValidParams("ID", "Username", "Email", "Password", "EmailConfirmed", "Joined", "LastLogin", "LastActive", ),
	)
    Router.RegisterApiEndpoint("User/Get", getUser, "GET",
        Router.MiddlewareValidParams("id"),
    )
	Router.RegisterApiEndpoint("User/Update", updateUser, "PATCH",
		Router.MiddlewareValidParams("id", "ID", "Username", "Email", "Password", "EmailConfirmed", "Joined", "LastLogin", "LastActive", ),
	)
	Router.RegisterApiEndpoint("User/Delete", deleteUser, "DELETE",
		Router.MiddlewareValidParams("id"),
	)
}