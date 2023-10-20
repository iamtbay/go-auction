package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/database"
	"github.com/iamtbay/go-auction/helpers"
	"github.com/iamtbay/go-auction/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auction struct{}

var auctionDB = database.AuctionDBInit()

// NEW AUCTION
// @Summary Start new auction
// @Description New auction for product
// @ID new-auction
// @Accept json
// @Produce json
// @Param userInfo body models.NewAuction true "New Auction Data"
// @Success 200 {object} map[string]interface{} "Auction succesfully created"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /auction/new [POST]
// @Tags auction
func (x *Auction) NewAuction(c *fiber.Ctx) error {
	var auctionInfo *models.NewAuction
	//get auction info
	err := c.BodyParser(&auctionInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//get current user id
	claims, err := helpers.GetUserClaimsFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//send it do db
	err = auctionDB.New(claims.ID, auctionInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Auction succesfully created",
	})
}

// GET AUCTION
// @Summary Get Auction
// @Description Get a auction
// @ID get-auction
// @Accept json
// @Produce json
// @Param id path string true "Get Auction Data"
// @Success 200 {object} map[string]interface{} "Succesful"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /auction/{id} [GET]
// @Tags auction
func (x *Auction) Get(c *fiber.Ctx) error {
	//Get params
	idString := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//get infos from db
	auctionInfo, err := auctionDB.Get(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//return
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succesful",
		"data":    auctionInfo,
	})
}
