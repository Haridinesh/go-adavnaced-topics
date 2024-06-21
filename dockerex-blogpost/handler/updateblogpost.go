package handler

import (
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UpdateBlogPost(c *fiber.Ctx) error {
	updatepost := models.Posts{}
	if err := c.BodyParser(&updatepost); err != nil {
		logger.Logging().Error("Cannot parse request body")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse request body", "Error": err, "Code": http.StatusInternalServerError})
	}
	idparam, err := c.ParamsInt("id")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": "Cannot parse query param"})
	}
	if err = repository.Dbn.ToUpdatePost(updatepost, idparam); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot update the record", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("updated post sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": nil, "Message": "Updated post sucessfully", "Error": nil, "Code": http.StatusOK})

}
