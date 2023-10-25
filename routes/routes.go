package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/iamtbay/go-auction/controllers/docs"
)

func SetupRoutes(app *fiber.App) {
	router := app.Group("/api/v1")

	router.Get("/swagger/*", swagger.HandlerDefault) //def

	AuthRoutes(router)
	ProductRoutes(router)
	AuctionRoutes(router)
	OfferRoutes(router)
}
