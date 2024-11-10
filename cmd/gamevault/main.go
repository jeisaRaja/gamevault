package main

import (
	"flag"
	"fmt"
	_ "gamevault/docs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

// @title GameVault API
// @version 1
// @description sample desc

// @contact.name Jeisa Raja
// @contact.url https://github.com/jeisaRaja

// @host localhost:8000
// @BasePath /v1
func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	r := chi.NewRouter()
	app.routes(r)

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", cfg.port))))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
