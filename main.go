package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/superg3m/stoic-go/Core"
)

func helloWorld(request *Core.StoicRequest, response Core.StoicResponse) {
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
	Core.LogDebug("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		Core.LogFatal(fmt.Sprintf("Server shutdown failed: %s", err))
	}
	Core.LogDebug("Server gracefully stopped.")
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

	//Core.RegisterPrefix("Api/0.1")
	Core.RegisterApiEndpoint("/User/Create", helloWorld, "POST")

	server := &http.Server{
		Addr:    SERVER_PORT,
		Handler: Core.Router,
	}

	//dsn := Core.GetDSN(DB_ENGINE, HOST, PORT, USER, PASSWORD, DBNAME)
	//db := Core.ConnectToDatabase(DB_ENGINE, dsn)
	//defer db.Close()

	siteSettings := Core.GetSiteSettings()
	fmt.Println(siteSettings["settings"].(map[string]any)["dbHost"])

	go gracefulShutdown(server)

	Core.LogDebug(fmt.Sprintf("Starting server on %s", SERVER_PORT))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
