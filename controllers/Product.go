package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/database"
	"github.com/iamtbay/go-auction/models"
)

type Product struct{}

var productDB = database.ProductDBInit()

// get
func (x *Product) Get(c *fiber.Ctx) error {
	slug := c.Params("slug")
	return c.SendString(fmt.Sprintf("hello get: %v product", slug))
}

// new
func (x *Product) New(c *fiber.Ctx) error {
	//get product infos from body
	var productInfo models.NewProductInfo
	err := c.BodyParser(&productInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//send it to db
	productDB.New(&productInfo)

	//return message
	return c.SendString("hello new product")
}

// update
func (x *Product) Update(c *fiber.Ctx) error {
	//get updated product infos from body id include.

	//send it to db

	//return message
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("hello update %v", id))
}

// delete
func (x *Product) Delete(c *fiber.Ctx) error {
	//get updated product infos from body id include.

	//send it to db

	//return msg
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("hello delete %v", id))
}
