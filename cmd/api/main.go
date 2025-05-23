package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// version
const version = "1.0.0"

// config of app
type config struct {
	port int
	env  string
}

// applicaion
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		config: cfg,
		logger: logger,
	}

	//http server which listens on the port provided in the config structure
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting server", "addr", srv.Addr, "env", cfg.env)
	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
