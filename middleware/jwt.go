package middleware

import (
	"cc-auth-service/helper"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return helper.Response(c, 401, "Unauthorized", nil)
		}
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "RajaBangkit" {
			return helper.Response(c, 401, "Unauthorized", nil)
		}
		token := parts[1]

		claims, err := helper.VerifyToken(token)
		if err != nil {
			return helper.Response(c, 401, "Unauthorized", nil)
		}
		fmt.Println(claims.Username)

		return c.Next()
	}
}
