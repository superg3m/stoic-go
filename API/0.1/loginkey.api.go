package API

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/LoginKey"
)

func createLoginKey(request *Router.StoicRequest, response Router.StoicResponse) {
	// Instantiate a new model object
	entity := LoginKey.New()
	entity.UserID = request.GetIntParam("UserID")
	entity.Provider = LoginKey.LoginKeyProvider(request.GetIntParam("Provider"))

	create := entity.Create()
	if create.IsBad() {
		response.SetError("Failed to create LoginKey | %s", create.GetError())
		return
	}

	response.SetData(fmt.Sprintf("LoginKey created successfully with ID %d", entity.UserID))
}

func getLoginKey(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")
	provider := request.GetIntParam("Provider")

	entity, err := LoginKey.FromUserID_Provider(id, provider)
	if err != nil {
		response.SetError(fmt.Sprintf("Error getting LoginKey: %s", err))
		return
	}

	response.SetData(entity)
}

func updateLoginKey(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")
	provider := request.GetIntParam("Provider")

	entity, err := LoginKey.FromUserID_Provider(id, provider)
	if err != nil {
		response.SetError("Failed to find LoginKey by ID | %s", err)
		return
	}
	entity.UserID = request.GetIntParam("UserID")
	entity.Provider = LoginKey.LoginKeyProvider(request.GetIntParam("Provider"))
	entity.Key = request.GetStringParam("Key")

	update := entity.Update()
	if update.IsBad() {
		response.SetError("Failed to update LoginKey | %s", update.GetError())
		return
	}

	response.SetData(fmt.Sprintf("LoginKey with ID %d updated successfully", id))
}

func deleteLoginKey(request *Router.StoicRequest, response Router.StoicResponse) {
	id := request.GetIntParam("id")
	provider := request.GetIntParam("provider")

	entity, err := LoginKey.FromUserID_Provider(id, provider)
	if err != nil {
		response.SetError("Failed to find LoginKey by ID | %s", err)
		return
	}

	delete := entity.Delete()

	if delete.IsBad() {
		response.SetError("Failed to delete LoginKey %s", delete.GetError())
		return
	}

	response.SetData(fmt.Sprintf("LoginKey with ID %d deleted successfully", id))
}

func init() {
	Router.RegisterApiEndpoint("LoginKey/Create", createLoginKey, "POST",
		Router.MiddlewareValidParams("id", "UserID", "Provider", "Key"),
	)
	Router.RegisterApiEndpoint("LoginKey/Get", getLoginKey, "GET",
		Router.MiddlewareValidParams("id"),
	)
	Router.RegisterApiEndpoint("LoginKey/Update", updateLoginKey, "PATCH",
		Router.MiddlewareValidParams("id", "UserID", "Provider"),
	)
	Router.RegisterApiEndpoint("LoginKey/Delete", deleteLoginKey, "DELETE",
		Router.MiddlewareValidParams("id"),
	)
}
