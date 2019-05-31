package main

import (
	"context"
	"os"
	"os/signal"
	"net/http"
	
	"github.com/Noah-Huppert/golog"
	"github.com/gorilla/mux"
)

func main() {
	// {{{1 Context
	ctx, ctxCancel := context.WithCancel(context.Background())

	// signals holds signals received by process
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals

		ctxCancel()
	}()

	// {{{1 Logger
	logger := golog.NewStdLogger("app-api")

	// {{{1 Configuration
	config, err := NewConfig()

	if err != nil {
		logger.Fatalf("failed to load configuration: %s", err.Error())
	}

	// {{{1 Router
	router := mux.NewRouter()

	// {{{1 Start HTTP server
	server := http.Server{
		Addr: config.HTTPAddr,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to serve: %s", err.Error())
		}
	}()

	logger.Infof("started server on %s", config.HTTPAddr)

	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Fatalf("failed to shutdown server: %s", err.Error())
	}

	logger.Info("done")
}
