package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnaucoin/govie/cmd/presenter"
	"github.com/tnaucoin/govie/internal/data"
	"strconv"
	"time"
)

func CreateMovieHandler(api *APIApp) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.SendString("create a new movie.")
		return nil
	}
}

func ShowMovieHandler(api *APIApp) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil || id < 1 {
			return api.notFoundErrorResponse(c)
		}
		m := &data.Movie{
			ID:        id,
			CreatedAt: time.Now(),
			Title:     "Ghostbusters",
			Year:      1984,
			Runtime:   105,
			Genres:    []string{"comedy", "adventure", "action"},
			Version:   1,
		}
		return c.JSON(presenter.MovieSuccessResponse(m))
	}
}
