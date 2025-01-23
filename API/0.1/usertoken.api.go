package API

import (
    "fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/UserToken"
)

func createUserToken(request *Router.StoicRequest, response Router.StoicResponse) {
	entity := UserToken.New()
    entity.ID = request.GetIntParam("ID")


    create := entity.Create()
    if create.IsBad() {
        response.SetError("Failed to create UserToken | %s", create.GetError())
        return
    }

	response.SetData(fmt.Sprintf("UserToken created successfully"))
}

func getUserToken(request *Router.StoicRequest, response Router.StoicResponse) {
    ID := request.GetIntParam("ID")

	entity, err := UserToken.FromID(ID)
	if err != nil {
		response.SetError("Failed to get UserToken | %s", err)
		return
	}

	response.SetData(entity)
}

func updateUserToken(request *Router.StoicRequest, response Router.StoicResponse) {
    ID := request.GetIntParam("ID")

	entity, err := UserToken.FromID(ID)
	if err != nil {
		response.SetError("Failed to get UserToken | %s", err)
		return
	}
    entity.ID = request.GetIntParam("ID")
    entity.UserID = request.GetIntParam("UserID")
    entity.Context = request.GetStringParam("Context")
    entity.Token = request.GetStringParam("Token")

	update := entity.Update()
	if update.IsBad() {
	    response.SetError("Failed to update UserToken | %s", update.GetError())
	    return
	}

	response.SetData(fmt.Sprintf("UserToken updated successfully"))
}

func deleteUserToken(request *Router.StoicRequest, response Router.StoicResponse) {
    ID := request.GetIntParam("ID")

	entity, err := UserToken.FromID(ID)
	if err != nil {
		response.SetError("Failed to get UserToken | %s", err)
		return
	}

	del := entity.Delete()

	if del.IsBad() {
	    response.SetError("Failed to delete UserToken %s", del.GetError())
	    return
	}

	response.SetData(fmt.Sprintf("UserToken deleted successfully"))
}

func init() {
	Router.RegisterApiEndpoint("UserToken/Create", createUserToken, "POST",
		Router.MiddlewareValidParams("ID", "UserID", "Created", "Context", "Token"),
	)
    Router.RegisterApiEndpoint("UserToken/Get", getUserToken, "GET",
        Router.MiddlewareValidParams("ID"),
    )
	Router.RegisterApiEndpoint("UserToken/Update", updateUserToken, "PATCH",
		Router.MiddlewareValidParams("ID", "UserID", "Created", "Context", "Token"),
	)
	Router.RegisterApiEndpoint("UserToken/Delete", deleteUserToken, "DELETE",
		Router.MiddlewareValidParams("ID"),
	)
}