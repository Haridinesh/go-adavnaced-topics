package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(c *fiber.Ctx) error {
	idparam, err := c.ParamsInt("id")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse query param", "Error": err, "Code": http.StatusBadRequest})
	}
	err = repository.Dbn.ToDeleteCategory(idparam)
	if err != nil {
		logger.Logging().Error("Cannot delete the record")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot delete the record", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Category deleted sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": nil, "Message": "Category deleted sucessfully", "Error": nil, "Code": http.StatusOK})

}
