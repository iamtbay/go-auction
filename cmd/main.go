package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtbay/go-auction/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file couldn't find")
	}
	app := fiber.New()

	//start db
	initDB()

	//setup routes
	routes.SetupRoutes(app)

	//start app
	app.Listen(":8080")
}
