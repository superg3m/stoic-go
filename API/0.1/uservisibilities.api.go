package API

import (
	"fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/UserVisibilities"
)

func createUserVisibilities(request *Router.StoicRequest, response Router.StoicResponse) {
	entity := UserVisibilities.New()
	entity.UserID = request.GetIntParam("UserID")

	create := entity.Create()
	if create.IsBad() {
		response.AddErrors(create.GetErrors(), "Failed to create UserVisibilities")
		return
	}

	response.SetData(fmt.Sprintf("UserVisibilities created successfully"))
}

func getUserVisibilities(request *Router.StoicRequest, response Router.StoicResponse) {
	UserID := request.GetIntParam("UserID")

	entity, errors := UserVisibilities.FromUserID(UserID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get UserVisibilities")
		return
	}

	response.SetData(entity)
}

func updateUserVisibilities(request *Router.StoicRequest, response Router.StoicResponse) {
	UserID := request.GetIntParam("UserID")

	entity, errors := UserVisibilities.FromUserID(UserID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get UserVisibilities from UserID")
		return
	}
	entity.UserID = request.GetIntParam("UserID")
	entity.Profile = request.GetBoolParam("Profile")
	entity.Email = request.GetBoolParam("Email")
	entity.Searches = request.GetBoolParam("Searches")
	entity.Birthday = request.GetBoolParam("Birthday")
	entity.RealName = request.GetBoolParam("RealName")
	entity.Description = request.GetBoolParam("Description")
	entity.Gender = request.GetBoolParam("Gender")

	update := entity.Update()
	if update.IsBad() {
		response.AddErrors(update.GetErrors(), "Failed to update UserVisibilities")
		return
	}

	response.SetData(fmt.Sprintf("UserVisibilities updated successfully"))
}

func deleteUserVisibilities(request *Router.StoicRequest, response Router.StoicResponse) {
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
		Router.MiddlewareValidParams("UserID", "Profile", "Email", "Searches", "Birthday", "RealName", "Description", "Gender"),
	)
	Router.RegisterApiEndpoint("UserVisibilities/Get", getUserVisibilities, "GET",
		Router.MiddlewareValidParams("UserID"),
	)
	Router.RegisterApiEndpoint("UserVisibilities/Update", updateUserVisibilities, "PATCH",
		Router.MiddlewareValidParams("UserID", "Profile", "Email", "Searches", "Birthday", "RealName", "Description", "Gender"),
	)
	Router.RegisterApiEndpoint("UserVisibilities/Delete", deleteUserVisibilities, "DELETE",
		Router.MiddlewareValidParams("UserID"),
	)
}
