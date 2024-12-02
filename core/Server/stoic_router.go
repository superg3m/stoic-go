package Server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}

func makeCompatible(handler StoicHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addCorsHeader(w)

		if r.Method == "OPTIONS" {
			return
		}

		stoicRequest := &StoicRequest{Request: r}
		stoicResponse := StoicResponse{ResponseWriter: w}

		handler(stoicRequest, stoicResponse)
	}
}

type StoicHandlerFunc func(r *StoicRequest, w StoicResponse)

var prefix string
var Router *mux.Router

func RegisterPrefix(newPrefix string) {
	prefix = newPrefix
}

func RegisterApiEndpoint(path string, functionEndpoint StoicHandlerFunc, method string) {
	resolvedPath := fmt.Sprintf("%s%s", prefix, path)
	Router.HandleFunc(resolvedPath, makeCompatible(functionEndpoint)).Methods(method, "OPTIONS")
}

func init() {
	Router = mux.NewRouter()
	prefix = ""
}
