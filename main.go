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

type StoicHandlerFunc func(r *Core.StoicRequest, w Core.StoicResponse)

// Enable CORS by adding headers
func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}

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

	Core.RegisterPrefix("Api/0.1/")
	mux := http.NewServeMux()
	mux.HandleFunc("/User/Create", makeCompatible(helloWorld))

	server := &http.Server{
		Addr:    SERVER_PORT,
		Handler: mux,
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
