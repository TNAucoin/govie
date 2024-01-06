package presenter

import "github.com/gofiber/fiber/v2"

func HealthCheckSuccessResponse(env string, version string) *fiber.Map {
	return &fiber.Map{
		"status": "available",
		"system_info": map[string]string{
			"environment": env,
			"version":     version,
		},
	}
}
