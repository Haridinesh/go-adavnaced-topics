package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetCommentsInPosts(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "cannot prase body param", "Error": err, "Code": http.StatusBadRequest})
	}
	data, err := repository.Dbn.ToGetCommentsByPostId(id)
	if err != nil {
		logger.Logging().Error("Cannot retrive data")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot retrive data", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Comments retrived sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": data, "Message": "comments present", "Error": nil, "Code": http.StatusOK})
}
