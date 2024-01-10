package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnaucoin/govie/internal/data"
	"github.com/tnaucoin/govie/internal/validator"
)

type Movie struct {
	ID      int64        `json:"id"`
	Title   string       `json:"title"`
	Year    int32        `json:"year,omitempty"`
	Runtime data.Runtime `json:"runtime,omitempty"`
	Genres  []string     `json:"genres,omitempty"`
}

func MovieSuccessResponse(data *data.Movie) *fiber.Map {
	movie := Movie{
		ID:      data.ID,
		Title:   data.Title,
		Year:    data.Year,
		Runtime: data.Runtime,
		Genres:  data.Genres,
	}
	return &fiber.Map{
		"movie": movie,
	}
}

func MovieValidationErrorResponse(err validator.ValidationError) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"movie":  "",
		"error":  err,
	}
}
