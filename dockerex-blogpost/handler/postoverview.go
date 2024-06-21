package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func OverviewOfPosts(c *fiber.Ctx) error {
	data, err := repository.Dbn.ToGetPostsOverview()
	if err != nil {
		logger.Logging().Error("cannot retrive the overview")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "cannot retrive the overview", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Overview for posts retrived")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": data, "Message": "Overview for posts retrived", "Error": nil, "Code": http.StatusOK})
}
