package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/controllers"
	"github.com/iamtbay/go-auction/middlewares"
)

func AuthRoutes(app fiber.Router) {
	var ctrllr = controllers.AuthInit()
	authR := app.Group("/auth")
	//
	authR.Post("/login", middlewares.LogoutMW, ctrllr.Login)
	authR.Post("/register", middlewares.LogoutMW, ctrllr.Register)
	authR.Patch("/update", middlewares.LoginMW, ctrllr.Update)
	authR.Post("/logout", middlewares.LoginMW, ctrllr.Logout)
}
