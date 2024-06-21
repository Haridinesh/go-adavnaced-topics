package handler

import (
	"blogpost/helper"
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateNewPost(c *fiber.Ctx) error {
	createpost := models.Posts{}
	err := c.BodyParser(&createpost)
	if err != nil {
		logger.Logging().Error("Cannot parse request body")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse request body", "Error": err, "Code": http.StatusInternalServerError})
	}
	createpost.Status = "Published"
	err = helper.StructValidation(createpost)
	if err != nil {
		logger.Logging().Error("Error occured in struct validation")
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error":   fmt.Sprintf("Error: %v", err),
			"message": "Resource not found",
			"code":    http.StatusNotFound,
		})
	}
	for _, v := range createpost.Categoriesid {
		err := repository.Dbn.ToCheckCategories(int(v))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "cateogry not found", "Error": err, "Code": http.StatusBadRequest})
		}
	}
	if err := repository.Dbn.ToCreatePost(createpost); err != nil {
		logger.Logging().Error("Cannot create task")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot create task", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Post created sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": "", "Message": "Post created sucessfully", "Error": nil, "Code": http.StatusOK})
}
