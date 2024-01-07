package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tnaucoin/govie/internal/validator"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type APIApp struct {
	config    config
	validator *validator.Validator
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
	v := validator.New()
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
