package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/controllers"
)

func AuthRoutes(app fiber.Router) {
	var ctrllr = controllers.AuthInit()
	authR := app.Group("/auth")
	//
	authR.Post("/login", ctrllr.Login)
	authR.Post("/register", ctrllr.Register)
	authR.Patch("/update", ctrllr.Update)
	authR.Post("/logout", ctrllr.Logout)
}
