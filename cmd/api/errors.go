package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
)

func (api *APIApp) logError(method, uri string, err error) {
	log.Errorw("", "method", method, "uri", uri, "error", err)
}

func errorResponse(status int, message string) error {
	return fiber.NewError(status, message)
}

func errorMessage(message string) *fiber.Map {
	return &fiber.Map{
		"error": message,
	}
}

func (api *APIApp) serverErrorResponse(c *fiber.Ctx, err error) error {
	api.logError(c.Method(), c.Request().URI().String(), err)
	return errorResponse(http.StatusInternalServerError,
		"the server encountered a problem and could not process your request",
	)
}

func (api *APIApp) notFoundErrorResponse() error {
	return errorResponse(http.StatusNotFound,
		"requested resource could not be found",
	)
}

func (api *APIApp) badRequestErrorResponse(c *fiber.Ctx, err error) error {
	return errorResponse(http.StatusBadRequest, err.Error())
}
