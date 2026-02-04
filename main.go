package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/superg3m/stoic-go/Core/ORM"

	_ "github.com/superg3m/stoic-go/API/0.1"
	_ "github.com/superg3m/stoic-go/Core/ORM"
	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/Core/Utility"
)

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	Utility.LogDebug("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		Utility.LogFatal(fmt.Sprintf("Server shutdown failed: %s", err))
	}
	Utility.LogDebug("Server gracefully stopped.")
}

func main() {
	const SERVER_PORT = ":8080"

	//Core.RegisterPrefix("API/0.1")

	server := &http.Server{
		Addr:    SERVER_PORT,
		Handler: Router.Router,
	}

	siteSettings := Utility.GetSiteSettings()
	siteSettings = siteSettings["settings"].(map[string]any)
	DB_ENGINE := Utility.CastAny[string](siteSettings["dbEngine"])
	HOST := Utility.CastAny[string](siteSettings["dbHost"])
	PORT := Utility.CastAny[int](siteSettings["dbPort"])
	USER := Utility.CastAny[string](siteSettings["dbUser"])
	PASSWORD := Utility.CastAny[string](siteSettings["dbPass"])
	DBNAMES := Utility.CastAny[[]string](siteSettings["dbNames"])

	for _, DBNAME := range DBNAMES {
		dsn := ORM.GetDSN(DB_ENGINE, HOST, PORT, USER, PASSWORD, DBNAME)
		ORM.Register(DBNAME, DB_ENGINE, dsn)
	}

	go gracefulShutdown(server)

	Utility.LogDebug(fmt.Sprintf("Starting server on %s", SERVER_PORT))
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed: %v", err)
	}

	for _, DBNAME := range DBNAMES {
		ORM.Close(DBNAME)
	}
}
