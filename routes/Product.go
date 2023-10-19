package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/controllers"
)

func ProductRoutes(app fiber.Router) {
	productR := app.Group("/product")
	ctrllr := controllers.ProductInit()

	//
	productR.Get("/:slug", ctrllr.Get)
	productR.Post("/new", ctrllr.New)
	productR.Patch("/update", ctrllr.Update)
	productR.Delete("/delete", ctrllr.Delete)

}
