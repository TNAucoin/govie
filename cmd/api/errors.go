package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
)

func (api *APIApp) logError(method, uri string, err error) {
	log.Errorw("", "method", method, "uri", uri, "error", err)
}

func errorResponse(c *fiber.Ctx, status int, m *fiber.Map) error {
	return c.Status(status).JSON(m)
}

func errorMessage(message string) *fiber.Map {
	return &fiber.Map{
		"error": message,
	}
}

func (api *APIApp) serverErrorResponse(c *fiber.Ctx, err error) error {
	api.logError(c.Method(), c.Request().URI().String(), err)
	return errorResponse(c, http.StatusInternalServerError, errorMessage(
		"the server encountered a problem and could not process your request",
	))
}

func (api *APIApp) notFoundErrorResponse(c *fiber.Ctx) error {
	return errorResponse(c, http.StatusNotFound, errorMessage(
		"requested resource could not be found",
	))
}
