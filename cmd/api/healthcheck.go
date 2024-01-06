package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnaucoin/govie/cmd/presenter"
)

func HealthcheckHandler(api *APIApp) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.JSON(presenter.HealthCheckSuccessResponse(api.config.env, version)); err != nil {
			return err
		}
		return nil
	}
}
