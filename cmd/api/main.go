package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type APIApp struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	apiApp := &APIApp{
		config: cfg,
	}

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(logger.New())

	apiApp.registerRoutes(app)

	log.Infow("", "port", cfg.port, "env", cfg.env)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.port)))
}
