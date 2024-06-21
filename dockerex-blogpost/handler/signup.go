package handler

import (
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ToSignUpUser(c *fiber.Ctx) error {
	signup := new(models.Logincredentials)
	err := c.BodyParser(signup)
	if err != nil {
		logger.Logging().Error("Failed reading the request body")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Data": nil, "Message": "Failed reading the request body", "Error": err, "Code": http.StatusInternalServerError})
	}
	if err := repository.Dbn.ToSignUpUser(signup); err != nil {
		logger.Logging().Error("Cannot Create the user")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot Create the user", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("User created sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": nil, "Message": "User created sucessfully", "Error": nil, "Code": http.StatusOK})
}
