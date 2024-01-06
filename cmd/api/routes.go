package main

import "github.com/gofiber/fiber/v2"

func (apiApp *APIApp) registerRoutes(app fiber.Router) {
	v1 := app.Group("v1")
	v1.Get("/healthcheck", HealthcheckHandler(apiApp))
	v1.Post("/movies", CreateMovieHandler(apiApp))
	v1.Get("/movies/:id", ShowMovieHandler(apiApp))
}
