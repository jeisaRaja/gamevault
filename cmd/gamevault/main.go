package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"gamevault/internal/data"
	"log"
	"net/http"
	"os"
	"time"

  _ "gamevault/docs"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
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
	var defaultPort = 8000
	var defaultDSN = "postgres://gamevault:password@localhost/gamevault"

	if portEnv := os.Getenv("APP_PORT"); portEnv != "" {
		fmt.Sscanf(portEnv, "%d", &defaultPort)
	}

	if DsnEnv := os.Getenv("DSN"); DsnEnv != "" {
		fmt.Sscanf(DsnEnv, "%s", &defaultDSN)
	}

	flag.IntVar(&cfg.port, "port", defaultPort, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", defaultDSN, "DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "10m", "PostgreSQL max idle connections")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("dsn is: ", cfg.db.dsn)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal("failed to connect to db, ", err)
	}
	defer db.Close()

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModel(db),
	}

	r := chi.NewRouter()
	app.routes(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
