package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetPostsByArchives(c *fiber.Ctx) error {
	startDate := c.Query("startdate")
	startDateTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return err
	}
	month := startDateTime.Month()
	year := startDateTime.Year()

	data, err := repository.Dbn.ToGetPostsByArchieves(year, month)
	if err != nil {
		logger.Logging().Error("Cannot get the posts by archives")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": data, "Message": "Cannot get posts by archives", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Posts by the archives retrived")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": data, "Message": "Posts by the archives retrived", "Error": nil, "Code": http.StatusOK})
}
