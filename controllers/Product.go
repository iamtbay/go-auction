package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Product struct{}

func (x *Product) Get(c *fiber.Ctx) error {
	slug := c.Params("slug")
	return c.SendString(fmt.Sprintf("hello get: %v product", slug))
}
func (x *Product) New(c *fiber.Ctx) error {
	return c.SendString("hello new product")
}
func (x *Product) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("hello update %v", id))
}
func (x *Product) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("hello delete %v", id))
}
