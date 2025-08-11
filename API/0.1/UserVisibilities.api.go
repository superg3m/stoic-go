package API

import (
    "fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/UserVisibilities"
)

func createUserVisibilities(request *Router.StoicRequest, response *Router.StoicResponse) {
	entity := UserVisibilities.New()
    entity.UserID = request.GetIntParam("UserID")
    entity.RealName = request.GetStringParam("RealName")
    entity.Description = request.GetStringParam("Description")
    entity.Gender = request.GetStringParam("Gender")


    create := entity.Create()
    if create.IsBad() {
        response.AddErrors(create.GetErrors(), "Failed to create UserVisibilities")
        return
    }

	response.SetData(fmt.Sprintf("UserVisibilities created successfully"))
}

func getUserVisibilities(request *Router.StoicRequest, response *Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")

	entity, errors := UserVisibilities.FromUserID(UserID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get UserVisibilities")
		return
	}

	response.SetData(entity)
}

func updateUserVisibilities(request *Router.StoicRequest, response *Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")

	entity, errors := UserVisibilities.FromUserID(UserID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get UserVisibilities")
		return
	}
    entity.UserID = request.GetIntParam("UserID")
    entity.RealName = request.GetStringParam("RealName")
    entity.Description = request.GetStringParam("Description")
    entity.Gender = request.GetStringParam("Gender")

	update := entity.Update()
	if update.IsBad() {
    	response.AddErrors(update.GetErrors(), "Failed to update UserVisibilities")
	    return
	}

	response.SetData(fmt.Sprintf("UserVisibilities updated successfully"))
}

func deleteUserVisibilities(request *Router.StoicRequest, response *Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")

	entity, errors := UserVisibilities.FromUserID(UserID)
	if errors != nil {
	    response.AddErrors(errors, "Failed to get UserVisibilities")
		return
	}

	del := entity.Delete()

	if del.IsBad() {
	    response.AddErrors(del.GetErrors(), "Failed to delete UserVisibilities")
	    return
	}

	response.SetData(fmt.Sprintf("UserVisibilities deleted successfully"))
}

func init() {
	Router.RegisterApiEndpoint("UserVisibilities/Create", createUserVisibilities, "POST",
		Router.MiddlewareValidParams("UserID", "RealName", "Description", "Gender"),
	)
    Router.RegisterApiEndpoint("UserVisibilities/Get", getUserVisibilities, "POST",
        Router.MiddlewareValidParams("UserID"),
    )
	Router.RegisterApiEndpoint("UserVisibilities/Update", updateUserVisibilities, "POST",
		Router.MiddlewareValidParams("UserID", "RealName", "Description", "Gender"),
	)
	Router.RegisterApiEndpoint("UserVisibilities/Delete", deleteUserVisibilities, "POST",
		Router.MiddlewareValidParams("UserID"),
	)
}