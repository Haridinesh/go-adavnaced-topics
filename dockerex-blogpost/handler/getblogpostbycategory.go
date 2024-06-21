package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"blogpost/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ToGetPostByCategory(c *fiber.Ctx) error {
	category := c.Query("category")

	categoriesid, err := repository.Dbn.ToGetCategoryByName(category)

	if err != nil {
		logger.Logging().Error("Error occured")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot get posts", "Error": nil, "Code": http.StatusBadRequest})
	}
	posts, err := repository.Dbn.ToGetPostByCategory(categoriesid)
	if err != nil {
		return err
	}
	postresp, err := service.Toblogpost(posts)
	if err != nil {
		return err
	}
	logger.Logging().Info("Posts fetched by the category")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": postresp, "Message": "Posts in the category", "Error": nil, "Code": http.StatusOK})
}
