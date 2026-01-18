package API

import (
	"fmt"

	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/TodoItem"
)

func createTodoItem(request *Router.StoicRequest, response *Router.StoicResponse) {
	entity := TodoItem.New()
	entity.OwnerID = request.GetIntParam("OwnerID")
	entity.Message = request.GetStringParam("Message")
	entity.Status = request.GetIntParam("Status")

	create := entity.Create()
	if create.IsBad() {
		response.AddErrors(create.GetErrors(), "Failed to create TodoItem")
		return
	}

	response.SetData(entity)
}

func getTodoItem(request *Router.StoicRequest, response *Router.StoicResponse) {
	OwnerID := request.GetIntParam("OwnerID")
	todos, err := TodoItem.AllFromOwnerID(OwnerID)
	if err != nil {
		response.AddError("Failed to get todo items")
		return
	}

	response.SetData(todos)
}

func updateTodoItem(request *Router.StoicRequest, response *Router.StoicResponse) {
	ID := request.GetIntParam("ID")

	entity, errors := TodoItem.FromID(ID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get TodoItem")
		return
	}

	entity.OwnerID = request.GetIntParam("OwnerID")
	entity.Message = request.GetStringParam("Message")
	entity.Status = request.GetIntParam("Status")

	update := entity.Update()
	if update.IsBad() {
		response.AddErrors(update.GetErrors(), "Failed to update TodoItem")
		return
	}

	response.SetData(fmt.Sprintf("TodoItem updated successfully"))
}

func deleteTodoItem(request *Router.StoicRequest, response *Router.StoicResponse) {
	ID := request.GetIntParam("ID")

	entity, errors := TodoItem.FromID(ID)
	if errors != nil {
		response.AddErrors(errors, "Failed to get TodoItem")
		return
	}

	del := entity.Delete()

	if del.IsBad() {
		response.AddErrors(del.GetErrors(), "Failed to delete TodoItem")
		return
	}

	response.SetData(fmt.Sprintf("TodoItem deleted successfully"))
}

func init() {
	Router.RegisterApiEndpoint("TodoItem", createTodoItem, "POST",
		Router.MiddlewareValidParams("OwnerID", "Message", "Status"),
	)
	Router.RegisterApiEndpoint("TodoItem", getTodoItem, "GET",
		Router.MiddlewareValidParams("OwnerID"),
	)
	Router.RegisterApiEndpoint("TodoItem", updateTodoItem, "PATCH",
		Router.MiddlewareValidParams("ID", "OwnerID", "Message", "Status"),
	)
	Router.RegisterApiEndpoint("TodoItem", deleteTodoItem, "DELETE",
		Router.MiddlewareValidParams("ID"),
	)
}
