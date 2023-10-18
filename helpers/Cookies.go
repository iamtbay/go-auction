package helpers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCookie(c *fiber.Ctx, token string) error {
	c.Cookie(&fiber.Cookie{
		Name:    "accessToken",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 12),
	})
	return nil
}

func DeleteCookie(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "accessToken",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})
	return nil
}
