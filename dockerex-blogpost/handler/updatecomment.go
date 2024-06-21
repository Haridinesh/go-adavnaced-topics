package handler

import (
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ToupdateComment(c *fiber.Ctx) error {
	updatecomment := models.Comments{}
	if err := c.BodyParser(&updatecomment); err != nil {
		logger.Logging().Error("Cannot parse request body")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse request body", "Error": err, "Code": http.StatusInternalServerError})
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse query param", "Error": err, "Code": http.StatusBadRequest})
	}
	if err = repository.Dbn.ToUpdateComment(updatecomment, id); err != nil {
		logger.Logging().Error("Cannot update comment")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot update comment", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Comment updated sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": "", "Message": "comment updated sucessfully", "Error": nil, "Code": http.StatusOK})
}
