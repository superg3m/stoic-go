package Router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/superg3m/stoic-go/Core/Utility"
)

type StoicHandlerFunc func(r *StoicRequest, w *StoicResponse)

var prefix string
var Router *mux.Router
var commonMiddlewares []StoicMiddleware

func init() {
	Router = mux.NewRouter()
	prefix = ""
	commonMiddlewares = nil
	MiddlewareRegisterCommon(MiddlewareCORS())
}

func RegisterPrefix(newPrefix string) {
	prefix = newPrefix
}

func adaptHandler(handler StoicHandlerFunc, middlewareList []StoicMiddleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stoicRequest := &StoicRequest{Request: r}
		stoicResponse := &StoicResponse{ResponseWriter: w}

		finalHandler := chainMiddleware(handler, middlewareList)
		finalHandler(stoicRequest, stoicResponse)

		if stoicResponse.GetErrors() != nil {
			stoicResponse.WriteHeader(http.StatusInternalServerError)
		}

		stoicResponse.Header().Set("Content-Type", "application/json")
		if stoicResponse.GetErrors() != nil {
			errorsJSON, err := json.Marshal(map[string][]string{"Errors": stoicResponse.GetErrors()})
			if err != nil {
				http.Error(stoicResponse, "Failed to marshal errors", http.StatusInternalServerError)
				return
			}

			_, err = fmt.Fprintf(stoicResponse, "%s", errorsJSON)
			Utility.AssertOnError(err)
		}
	}
}

func RegisterApiEndpoint(path string, handler StoicHandlerFunc, method string, middlewares ...StoicMiddleware) {
	Utility.Assert(commonMiddlewares != nil)

	resolvedPath := fmt.Sprintf("%s/%s", prefix, path)
	middlewareList := append(commonMiddlewares, middlewares...)

	Router.HandleFunc(resolvedPath, adaptHandler(handler, middlewareList)).Methods(method, "OPTIONS")
}
