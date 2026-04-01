package controllers

import (
	"lms/database"
	"lms/models"

	"github.com/gofiber/fiber/v3"
)

func RegisterSiswa(c fiber.Ctx) error {
	var input models.Student

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "format tidak valid",
		})
	}
	result := database.DB.Create(&input)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal mendaftarkan siswa",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "berhasil daftar",
		"data":    input,
	})
}
