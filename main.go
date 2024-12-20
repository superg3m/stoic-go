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

	_ "github.com/superg3m/stoic-go/API/0.1"
	"github.com/superg3m/stoic-go/Core/ORM"
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

const (
	DB_ENGINE = "mysql"
	HOST      = "localhost"
	PORT      = 3306
	USER      = "root"
	PASSWORD  = "P@55word"
	DBNAME    = "stoic"
)

func main() {
	const SERVER_PORT = ":8080"

	//Core.RegisterPrefix("API/0.1")

	server := &http.Server{
		Addr:    SERVER_PORT,
		Handler: Router.Router,
	}

	dsn := ORM.GetDSN(DB_ENGINE, HOST, PORT, USER, PASSWORD, DBNAME)
	ORM.Connect(DB_ENGINE, dsn)
	defer ORM.Close()

	go gracefulShutdown(server)

	Utility.LogDebug(fmt.Sprintf("Starting server on %s", SERVER_PORT))
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed: %v", err)
	}
}
