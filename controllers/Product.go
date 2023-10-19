package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/database"
	"github.com/iamtbay/go-auction/helpers"
	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{}

var productDB = database.ProductDBInit()

// get
func (x *Product) Get(c *fiber.Ctx) error {
	slug := c.Params("slug")
	//
	product, err := productDB.Get(slug)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succesful",
		"data":    product,
	})
}

// new
func (x *Product) New(c *fiber.Ctx) error {
	//get product infos from body
	var productInfo *models.NewProductInfo
	err := c.BodyParser(&productInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	claims, err := helpers.GetUserClaimsFromJWT(c)
	if err != nil {
		return err
	}
	productInfo.SellerID = claims.ID
	productInfo.CreatedAt = time.Now()
	//send it to db
	err = productDB.New(productInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//return message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created Succesfully",
	})
}

// update
func (x *Product) Update(c *fiber.Ctx) error {
	//get updated product infos from body id include.
	var productInfo *models.GetProductInfo
	err := c.BodyParser(&productInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	claims, err := helpers.GetUserClaimsFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//send it to db
	err = productDB.Update(claims.ID, productInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//return message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated!",
	})
}

// delete
func (x *Product) Delete(c *fiber.Ctx) error {
	var Product struct {
		ID string `json:"_id" bson:"_id"`
	}
	err := c.BodyParser(&Product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//get updated product infos from body id include.
	claims, err := helpers.GetUserClaimsFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//send it to db
	obId, err := primitive.ObjectIDFromHex(Product.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = productDB.Delete(claims.ID, obId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//return msg

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%v deleted succesfully", obId),
	})
}
