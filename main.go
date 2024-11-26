package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/superg3m/stoic-go/Core"
)

type StoicHandlerFunc func(r *Core.StoicRequest, w Core.StoicResponse)

// Enable CORS by adding headers
func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
}

func helloWorld(request *Core.StoicRequest, response Core.StoicResponse) {
	//request.PrintRequestData()

	if !request.HasAll("username", "email") {
		fmt.Fprintf(response, "Invalid Params")
		return
	}

	username, _ := request.GetStringParam("username")

	fmt.Fprintf(response, "Hello %s", string(username))
}

// makeCompatible adapts StoicHandlerFunc to http.HandlerFunc
func makeCompatible(handler StoicHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addCorsHeader(w)

		if r.Method == "OPTIONS" {
			return
		}

		stoicRequest := &Core.StoicRequest{Request: r}
		stoicResponse := Core.StoicResponse{ResponseWriter: w}

		handler(stoicRequest, stoicResponse)
	}
}

func main() {
	const SERVER_PORT = ":8080"

	Core.RegisterPrefix("Api/0.1/")
	http.HandleFunc("/User/Create", makeCompatible(helloWorld))

	fmt.Println("Cors Fix Server started on Port:", SERVER_PORT)
	log.Fatal(http.ListenAndServe(SERVER_PORT, nil))
}
