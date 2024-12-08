package Server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/superg3m/stoic-go/core/Client"
	"net/http"
)

type StoicHandlerFunc func(r *Client.StoicRequest, w StoicResponse)

var prefix string
var Router *mux.Router
var commonMiddlewares []StoicMiddleware

func init() {
	Router = mux.NewRouter()
	prefix = ""
	commonMiddlewares = []StoicMiddleware{}
}

func makeCompatible(handler StoicHandlerFunc, middlewareList []StoicMiddleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stoicRequest := &Client.StoicRequest{Request: r}
		stoicResponse := StoicResponse{ResponseWriter: w}

		finalHandler := chainMiddleware(handler, middlewareList)
		finalHandler(stoicRequest, stoicResponse)
	}
}

func RegisterPrefix(newPrefix string) {
	prefix = newPrefix
}

func RegisterApiEndpoint(path string, handler StoicHandlerFunc, method string, middlewares ...StoicMiddleware) {
	middlewareList := append(commonMiddlewares, middlewares...)

	resolvedPath := fmt.Sprintf("%s%s", prefix, path)
	Router.HandleFunc(resolvedPath, makeCompatible(handler, middlewareList)).Methods(method, "OPTIONS")
}
