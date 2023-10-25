package controllers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/iamtbay/go-auction/database"
	"github.com/iamtbay/go-auction/helpers"
	"github.com/iamtbay/go-auction/models"
)

type Offers struct{}

var offerDB = database.OfferDBInit()
var clients = make(map[string]map[*websocket.Conn]bool)
var broadcast = make(chan *models.OfferInfo)

// @ID new-offer-auction
// @Summary New Offer For Auction
// @Description Post new bid for auction
// @Param newOfferInfo body models.OfferInfo true "New Offer Info"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Offer sent"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /offers/new [post]
// @Tags auction
func (x *Offers) NewOffer(c *fiber.Ctx) error {
	var newOffer *models.OfferInfo
	err := c.BodyParser(&newOffer)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//get user id
	claims, err := helpers.GetUserClaimsFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	newOffer.UserID = claims.ID
	newOffer.CreatedAt = time.Now()
	//send vars to db
	err = offerDB.NewOffer(newOffer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//send to channel
	broadcast <- newOffer
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "success",
	})
}

// @ID get-ws-auction
// @Summary Get Auction
// @Description Get Real-Time Bids for auction
// @Param auctionId path string true "Get Auction"
// @Produce json
// @Success 200 {object} models.OfferInfo "Offer got"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /offers/{id} [get]
// @Tags auction
func (x *Offers) GetAuction(c *websocket.Conn) {
	//todo it will get auction details by id.
	//add user to clients
	auctionID := c.Params("id")
	if _, exists := clients[auctionID]; !exists {
		clients[auctionID] = make(map[*websocket.Conn]bool)
	}
	clients[auctionID][c] = true
	defer func() {
		delete(clients[auctionID], c)
		c.Close()
	}()

	for conn := range clients[auctionID] {
		for offers := range broadcast {
			if err := conn.WriteJSON(offers); err != nil {
				log.Println("off err")
				return
			}
		}

	}

}

