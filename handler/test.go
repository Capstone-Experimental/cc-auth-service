package handler

import (
	"cc-auth-service/helper"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	return helper.Response(c, 200, "Success to GET test", nil)
}

func Dashboard(c *fiber.Ctx) error {
	return helper.Response(c, 200, "Success to GET dashboard", nil)
}
