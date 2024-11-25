package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/superg3m/stoic-go/Core"
)

type StoicHandlerFunc func(r *Core.StoicRequest, w Core.StoicResponse)

func helloWorld(request *Core.StoicRequest, response Core.StoicResponse) {
	fmt.Fprintf(response, "Hello world")
}

// makeCompatible adapts StoicHandlerFunc to http.HandlerFunc
func makeCompatible(handler StoicHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stoicRequest := &Core.StoicRequest{Request: r}
		stoicResponse := Core.StoicResponse{w}

		handler(stoicRequest, stoicResponse)
	}
}

func main() {
	const SERVER_PORT = ":8080"
	// Database credentials

	Core.RegisterPrefix("Api/0.1/")
	http.HandleFunc("/Hello", makeCompatible(helloWorld))

	// start http server
	fmt.Println("Server started on Port:", SERVER_PORT)
	log.Fatal(http.ListenAndServe(SERVER_PORT, nil))
}
