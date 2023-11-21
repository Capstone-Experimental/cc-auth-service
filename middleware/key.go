package middleware

import (
	"cc-auth-service/helper"

	"github.com/gofiber/fiber/v2"
)

var allowedAPIKeys = map[string]bool{
	"140jutaInkubasi": true,
}

func ApiKey() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		if _, ok := allowedAPIKeys[apiKey]; !ok {
			return helper.Response(c, 401, "Unauthorized", nil)
		}
		return c.Next()
	}
}
