package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/tnaucoin/govie/internal/validator"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type APIApp struct {
	config    config
	validator *validator.Validator
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://govie:pa55word@db/govie?sslmode=disable", "DB DSN")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "Psql max idle connections.")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "Psql max open connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "Psql max idle connection time")

	flag.Parse()
	v := validator.New()
	db, err := openDB(cfg)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer db.Close()
	log.Info("db connection pool established")

	apiApp := &APIApp{
		config:    cfg,
		validator: v,
	}

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			type EResp struct {
				Error string `json:"error"`
			}
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(code).JSON(&EResp{Error: err.Error()})
		},
	})

	app.Use(logger.New(), recover.New())

	apiApp.registerRoutes(app)

	log.Infow("", "port", cfg.port, "env", cfg.env)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.port)))
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
