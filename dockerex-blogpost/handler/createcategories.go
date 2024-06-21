package handler

import (
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreatingNewCategory(c *fiber.Ctx) error {
	category := models.Categories{}
	err := c.BodyParser(&category)
	if err != nil {
		logger.Logging().Error("Cannot parse request body")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse request body", "Error": err, "Code": http.StatusInternalServerError})
	}
	err = repository.Dbn.ToCreateCategory(category)
	if err != nil {
		logger.Logging().Error("Cannot create new category")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": "", "Message": "Cannot create new category", "Error": nil, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Category created sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": "", "Message": "category created sucessfully", "Error": nil, "Code": http.StatusOK})
}
