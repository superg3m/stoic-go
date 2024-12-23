package API

import (
    "fmt"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/UserRole"
)

func createUserRole(request *Router.StoicRequest, response Router.StoicResponse) {
	entity := UserRole.New()
    entity.UserID = request.GetIntParam("UserID")
    entity.RoleID = request.GetIntParam("RoleID")


    create := entity.Create()
    if create.IsBad() {
        response.SetError("Failed to create UserRole | %s", create.GetError())
        return
    }

	response.SetData(fmt.Sprintf("UserRole created successfully"))
}

func getUserRole(request *Router.StoicRequest, response Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")
    RoleID := request.GetIntParam("RoleID")

	entity, err := UserRole.FromUserID_RoleID(UserID, RoleID)
	if err != nil {
		response.SetError("Failed to get UserRole | %s", err)
		return
	}

	response.SetData(entity)
}

func updateUserRole(request *Router.StoicRequest, response Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")
    RoleID := request.GetIntParam("RoleID")

	entity, err := UserRole.FromUserID_RoleID(UserID, RoleID)
	if err != nil {
		response.SetError("Failed to get UserRole | %s", err)
		return
	}
    entity.UserID = request.GetIntParam("UserID")
    entity.RoleID = request.GetIntParam("RoleID")

	update := entity.Update()
	if update.IsBad() {
	    response.SetError("Failed to update UserRole | %s", update.GetError())
	    return
	}

	response.SetData(fmt.Sprintf("UserRole updated successfully"))
}

func deleteUserRole(request *Router.StoicRequest, response Router.StoicResponse) {
    UserID := request.GetIntParam("UserID")
    RoleID := request.GetIntParam("RoleID")

	entity, err := UserRole.FromUserID_RoleID(UserID, RoleID)
	if err != nil {
		response.SetError("Failed to get UserRole | %s", err)
		return
	}

	del := entity.Delete()

	if del.IsBad() {
	    response.SetError("Failed to delete UserRole %s", del.GetError())
	    return
	}

	response.SetData(fmt.Sprintf("UserRole deleted successfully"))
}

func init() {
	Router.RegisterApiEndpoint("UserRole/Create", createUserRole, "POST",
		Router.MiddlewareValidParams("UserID", "RoleID"),
	)
    Router.RegisterApiEndpoint("UserRole/Get", getUserRole, "GET",
        Router.MiddlewareValidParams("UserID", "RoleID"),
    )
	Router.RegisterApiEndpoint("UserRole/Update", updateUserRole, "PATCH",
		Router.MiddlewareValidParams("UserID", "RoleID"),
	)
	Router.RegisterApiEndpoint("UserRole/Delete", deleteUserRole, "DELETE",
		Router.MiddlewareValidParams("UserID", "RoleID"),
	)
}