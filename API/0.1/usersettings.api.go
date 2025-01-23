package API

import (
    "fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/UserSettings"
)

func createUserSettings(request *Router.StoicRequest, response Router.StoicResponse) {
	entity := UserSettings.New()
    entity.UserID = request.GetIntParam("UserID")


    create := entity.Create()
    if create.IsBad() {
        response.SetError("Failed to create UserSettings | %s", create.GetError())
        return
    }

	response.SetData(fmt.Sprintf("UserSettings created successfully"))
}

func getUserSettings(request *Router.StoicRequest, response Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")

	entity, err := UserSettings.FromUserID(UserID)
	if err != nil {
		response.SetError("Failed to get UserSettings | %s", err)
		return
	}

	response.SetData(entity)
}

func updateUserSettings(request *Router.StoicRequest, response Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")

	entity, err := UserSettings.FromUserID(UserID)
	if err != nil {
		response.SetError("Failed to get UserSettings | %s", err)
		return
	}
    entity.UserID = request.GetIntParam("UserID")
    entity.HtmlEmails = request.GetBoolParam("HtmlEmails")
    entity.PlaySounds = request.GetBoolParam("PlaySounds")

	update := entity.Update()
	if update.IsBad() {
	    response.SetError("Failed to update UserSettings | %s", update.GetError())
	    return
	}

	response.SetData(fmt.Sprintf("UserSettings updated successfully"))
}

func deleteUserSettings(request *Router.StoicRequest, response Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")

	entity, err := UserSettings.FromUserID(UserID)
	if err != nil {
		response.SetError("Failed to get UserSettings | %s", err)
		return
	}

	del := entity.Delete()

	if del.IsBad() {
	    response.SetError("Failed to delete UserSettings %s", del.GetError())
	    return
	}

	response.SetData(fmt.Sprintf("UserSettings deleted successfully"))
}

func init() {
	Router.RegisterApiEndpoint("UserSettings/Create", createUserSettings, "POST",
		Router.MiddlewareValidParams("UserID", "HtmlEmails", "PlaySounds"),
	)
    Router.RegisterApiEndpoint("UserSettings/Get", getUserSettings, "GET",
        Router.MiddlewareValidParams("UserID"),
    )
	Router.RegisterApiEndpoint("UserSettings/Update", updateUserSettings, "PATCH",
		Router.MiddlewareValidParams("UserID", "HtmlEmails", "PlaySounds"),
	)
	Router.RegisterApiEndpoint("UserSettings/Delete", deleteUserSettings, "DELETE",
		Router.MiddlewareValidParams("UserID"),
	)
}