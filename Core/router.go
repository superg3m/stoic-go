package Core

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type I_ApiEndpoint interface {
	RegisterApiEndpoints()
}

var prefix string
var router *mux.Router

func init() {
	router = mux.NewRouter()
}

func RegisterPrefix(newPrefix string) {
	prefix = newPrefix
}

func RegisterEndpoint(path string, functionEndpoint func(w http.ResponseWriter, r *http.Request), method string) {
	resolvedPath := fmt.Sprintf("%s/%s", prefix, path)
	router.HandleFunc(resolvedPath, functionEndpoint).Methods(method)
}
