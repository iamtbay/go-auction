package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/iamtbay/go-auction/controllers"
	"github.com/iamtbay/go-auction/middlewares"
)

func OfferRoutes(app fiber.Router) {
	controller := controllers.OffersInit()
	offerR := app.Group("/offers")

	offerR.Get("/:id", websocket.New(controller.GetAuction))
	offerR.Use(middlewares.LoginMW)
	offerR.Post("/new", controller.NewOffer)
}
