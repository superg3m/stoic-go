package Server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/superg3m/stoic-go/core/Client"
	"github.com/superg3m/stoic-go/core/Middleware"
	"net/http"
	"reflect"
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

// Wrap the StoicHandlerFunc with middleware for compatibility.
func makeCompatible(handler StoicHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		stoicRequest := &Client.StoicRequest{Request: r}
		stoicResponse := StoicResponse{ResponseWriter: w}

		handler(stoicRequest, stoicResponse)
	}
}

func RegisterCommonMiddleware(middlewares ...StoicMiddleware) {
	for _, middleware := range middlewares {
		// Check if middleware is already in the commonMiddlewares list (to avoid duplicates)
		if !isMiddlewareRegistered(middleware) {
			commonMiddlewares = append(commonMiddlewares, middleware)
		}
	}
}

// Check if a middleware is already registered
func isMiddlewareRegistered(middleware StoicMiddleware) bool {
	for _, existingMiddleware := range commonMiddlewares {
		// Compare functions by their reflect value to check if they are the same
		if reflect.DeepEqual(existingMiddleware, middleware) {
			return true
		}
	}
	return false
}

func RegisterPrefix(newPrefix string) {
	prefix = newPrefix
}

func chainMiddleware(handler StoicHandlerFunc, middlewares []Middleware.StoicMiddleware) StoicHandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func RegisterApiEndpoint(path string, handler StoicHandlerFunc, method string, middlewares ...StoicMiddleware) {
	middlewareList := append(commonMiddlewares, middlewares...)
	finalHandler := chainMiddleware(handler, middlewareList)

	resolvedPath := fmt.Sprintf("%s%s", prefix, path)
	Router.HandleFunc(resolvedPath, makeCompatible(finalHandler)).Methods(method, "OPTIONS")
}
