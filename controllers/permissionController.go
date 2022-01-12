package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matiolsz/go-and-react/database"
	"github.com/matiolsz/go-and-react/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Role

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
