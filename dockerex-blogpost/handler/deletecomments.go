package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteComment(c *fiber.Ctx) error {
	idparam, err := c.ParamsInt("id")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse query param", "Error": err, "Code": http.StatusBadRequest})
	}
	err = repository.Dbn.ToDeleteComments(uint64(idparam))
	if err != nil {
		logger.Logging().Error("Cannot delete the record")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot delete the record", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Comment deleted sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": nil, "Message": "comment Deleted sucessfully", "Error": nil, "Code": http.StatusOK})

}
