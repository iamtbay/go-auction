package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// LOGIN MW
func LoginMW(c *fiber.Ctx) error {
	token := c.Cookies("accessToken")
	if token != "" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Please Login First!",
	})
}

// LOGOUT MW
func LogoutMW(c *fiber.Ctx) error {
	token := c.Cookies("accessToken")
	if token == "" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Please Logout First!",
	})
}
