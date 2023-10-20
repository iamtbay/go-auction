package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/controllers"
	"github.com/iamtbay/go-auction/middlewares"
)

func AuctionRoutes(app fiber.Router) {
	controller := controllers.AuctionInit()
	auctionR := app.Group("/auction")
	//
	auctionR.Get("/:id", controller.Get)
	auctionR.Use(middlewares.LoginMW)
	auctionR.Post("/new", controller.NewAuction)
}
