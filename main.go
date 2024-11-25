package main

import (
	"log"
	"net/http"

	"github.com/superg3m/stoic-go/Core"
)

func main() {
	// Database credentials

	Core.RegisterPrefix("Api/0.1/")

	// start http server
	log.Fatal(http.ListenAndServe(SERVER_PORT, nil))
}
