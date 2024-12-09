package Server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/superg3m/stoic-go/core/Client"
)

type StoicHandlerFunc func(r *Client.StoicRequest, w StoicResponse)

var prefix string
var Router *gin.Engine
var commonMiddlewares []StoicMiddleware

func init() {
	gin.SetMode(gin.ReleaseMode)

	// Router = gin.Default() // Logger, Recovery Middleware
	Router = gin.New()
	prefix = ""
	commonMiddlewares = []StoicMiddleware{}
}

func RegisterPrefix(newPrefix string) {
	prefix = newPrefix
}

func adaptHandler(handler StoicHandlerFunc, middlewareList []StoicMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		stoicRequest := &Client.StoicRequest{Request: c.Request}
		stoicResponse := StoicResponse{ResponseWriter: c.Writer}

		finalHandler := chainMiddleware(handler, middlewareList)
		finalHandler(stoicRequest, stoicResponse)
	}
}

func RegisterApiEndpoint(path string, handler StoicHandlerFunc, method string, middlewares ...StoicMiddleware) {
	resolvedPath := fmt.Sprintf("%s%s", prefix, path)
	allMiddlewares := append(commonMiddlewares, middlewares...)

	switch method {
	case "GET":
		Router.GET(resolvedPath, adaptHandler(handler, allMiddlewares)).OPTIONS(resolvedPath, adaptHandler(handler, allMiddlewares))
	case "POST":
		Router.POST(resolvedPath, adaptHandler(handler, allMiddlewares)).OPTIONS(resolvedPath, adaptHandler(handler, allMiddlewares))
	case "PUT":
		Router.PUT(resolvedPath, adaptHandler(handler, allMiddlewares)).OPTIONS(resolvedPath, adaptHandler(handler, allMiddlewares))
	case "DELETE":
		Router.DELETE(resolvedPath, adaptHandler(handler, allMiddlewares)).OPTIONS(resolvedPath, adaptHandler(handler, allMiddlewares))
	default:
		panic("Unsupported method: " + method)
	}
}
