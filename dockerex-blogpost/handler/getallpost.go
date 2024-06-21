package handler

import (
	"blogpost/logger"
	"blogpost/repository"
	"blogpost/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllBlogPosts(c *fiber.Ctx) error {

	posts, err := repository.Dbn.ToGetAllPosts()

	if err != nil {
		logger.Logging().Error("Error occured")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot get posts", "Error": nil, "Code": http.StatusBadRequest})
	}
	postresp, err := service.Toblogpost(posts)
	if err != nil {
		return err
	}
	logger.Logging().Info("Fetched all posts")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": postresp, "Message": "Blogpost that are found", "Error": nil, "Code": http.StatusOK})
}
