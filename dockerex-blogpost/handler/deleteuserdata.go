package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteUserData(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse query param", "Error": err, "Code": http.StatusBadRequest})
	}
	if err := repository.Dbn.ToDeleteUserData(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "cannot delete user", "Error": err, "Code": http.StatusBadRequest})
	}
	err = repository.Dbn.ToDeleteRedisCache(c.Context().Value("uuid").(string))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot invalid user credentials", "Error": err, "Code": http.StatusBadRequest})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": "", "Message": "User deleted sucessfully", "Error": nil, "Code": http.StatusOK})
}
