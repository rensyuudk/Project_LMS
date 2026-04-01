package main

import (
	"lms/database"
	"lms/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	routes.SetUp(app)
	app.Listen(":8000")
}
