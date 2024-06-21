package handler

import (
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UpdateUserData(c *fiber.Ctx) error {
	userdata := models.Logincredentials{}

	if err := c.BodyParser(&userdata); err != nil {
		logger.Logging().Error("Cannot parse request body")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse request body", "Error": err, "Code": http.StatusBadRequest})
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse query param", "Error": err, "Code": http.StatusBadRequest})
	}
	err = repository.Dbn.ToUpdataUserData(userdata, id)
	if err != nil {
		logger.Logging().Error("Cannot update user record")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot update userdata", "Error": err, "Code": http.StatusBadRequest})
	}
	err = repository.Dbn.ToDeleteRedisCache(c.Context().Value("uuid").(string))
	if err != nil {
		logger.Logging().Error("Cannot delete in redis cache")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot invalid user credentials", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("User updated sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": nil, "Message": "User updated sucessfully", "Error": nil, "Code": http.StatusOK})
}
