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

// @Summary Get Product
// @Description Get Product Details
// @ID get-product
// @Produce json
// @Param slug path string true "Slug"
// @Success 200 {object} map[string]interface{} "User Get succesfully"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /product/{slug} [get]
// @Tags product
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

// @Summary New Product
// @Description Add a product
// @ID new-product
// @Accept json
// @Produce json
// @Param newProductInfo body models.NewProductInfo true "New product data"
// @Success 200 {object} map[string]interface{} "Product created succesfully"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /product/new [post]
// @Tags product
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

// @Summary Update Product
// @Description Update a product infos
// @ID update-product
// @Accept json
// @Produce json
// @Param getProductInfo body models.GetProductInfo true "Update product data"
// @Success 200 {object} map[string]interface{} "Product updated succesfully"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /product/update [patch]
// @Tags product
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

// @Summary Delete Product
// @Description Delete a product
// @ID delete-product
// @Accept json
// @Produce json
// @Param productID body models.DeleteProduct true "Delete a product"
// @Success 200 {object} map[string]interface{} "Product deleted succesfully"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /product/delete [delete]
// @Tags product
func (x *Product) Delete(c *fiber.Ctx) error {
	var Product *models.DeleteProduct
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
