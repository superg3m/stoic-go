package main

import (
	"context"
	"fmt"
	Server2 "github.com/superg3m/stoic-go/cmd/src/Server"
	Utility2 "github.com/superg3m/stoic-go/cmd/src/Utility"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func helloWorld(request *Server2.StoicRequest, response Server2.StoicResponse) {
	if !request.HasAll("username", "email") {
		response.SetError("Invalid Params")
		return
	}

	username := request.GetStringParam("username")

	if len(username) < 8 {
		response.SetError("Username must be at least 8 characters long")
		return
	}

	if username != "superg3m" {
		response.SetError("Wrong Password")
		return
	}

	fmt.Fprintf(response, "Hello %s", username)
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	Utility2.LogDebug("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		Utility2.LogFatal(fmt.Sprintf("Server shutdown failed: %s", err))
	}
	Utility2.LogDebug("Server gracefully stopped.")
}

const (
	DB_ENGINE = "mysql"
	HOST      = "localhost"
	PORT      = 5432
	USER      = "jacob"
	PASSWORD  = "password"
	DBNAME    = "bookstoreDB"
)

func main() {
	const SERVER_PORT = ":8080"

	//core.RegisterPrefix("api/0.1")
	Server2.RegisterApiEndpoint("/User/Create", helloWorld, "POST")

	server := &http.Server{
		Addr:    SERVER_PORT,
		Handler: Server2.Router,
	}

	//dsn := core.GetDSN(DB_ENGINE, HOST, PORT, USER, PASSWORD, DBNAME)
	//db := core.ConnectToDatabase(DB_ENGINE, dsn)
	//defer db.Close()

	siteSettings := Utility2.GetSiteSettings()
	fmt.Println(siteSettings["settings"].(map[string]any)["dbHost"])

	go gracefulShutdown(server)

	Utility2.LogDebug(fmt.Sprintf("Starting server on %s", SERVER_PORT))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
