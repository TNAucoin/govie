package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tnaucoin/govie/cmd/presenter"
	"github.com/tnaucoin/govie/internal/data"
	"strconv"
	"strings"
	"time"
)

func CreateMovieHandler(api *APIApp) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			Title   string   `json:"title" validate:"required"`
			Year    int32    `json:"year" validate:"required,gte=1888,lte=9999"`
			Runtime int32    `json:"runtime" validate:"required,gte=1"`
			Genres  []string `json:"genres" validate:"required,min=1"`
		}
		err := c.BodyParser(&input)
		if err != nil {
			return api.badRequestErrorResponse(err)
		}
		if errs := api.validator.Validate(input); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)

			for _, err := range errs {
				errMsgs = append(errMsgs, fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'",
					err.FailedField,
					err.Value,
					err.Tag,
				))
			}
			return &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: strings.Join(errMsgs, ", "),
			}
		}
		b := &data.Movie{
			Title:   input.Title,
			Year:    input.Year,
			Runtime: data.Runtime(input.Runtime),
			Genres:  input.Genres,
		}
		return c.JSON(presenter.MovieSuccessResponse(b))
	}
}

func ShowMovieHandler(api *APIApp) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil || id < 1 {
			return api.notFoundErrorResponse()
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
