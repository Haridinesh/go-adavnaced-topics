package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LogoutUser(c *fiber.Ctx) error {
	err := repository.Dbn.ToDeleteRedisCache(c.Context().Value("uuid").(string))
	if err != nil {
		logger.Logging().Error("Cannot delete rediscache")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot delete rediscache", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("User logger out sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": "", "Message": "logged out sucessfully", "Error": nil, "Code": http.StatusOK})
}
