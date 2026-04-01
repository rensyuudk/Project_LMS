package routes

import (
	"lms/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetUp(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterSiswa)
}
