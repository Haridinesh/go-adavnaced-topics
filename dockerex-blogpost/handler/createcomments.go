package handler

import (
	"blogpost/helper"
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"strconv"
	"strings"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddCommentToPost(c *fiber.Ctx) error {
	comment := models.Comments{}
	err := c.BodyParser(&comment)
	if err != nil {
		logger.Logging().Error("Cannot parse request body")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse request body", "Error": err, "Code": http.StatusInternalServerError})
	}
	token := c.Get(fiber.HeaderAuthorization)
	orgtoken, ok := strings.CutPrefix(token, "Bearer ")
	if !ok {
		logger.Logging().Error("Cannot parse tokenstring")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse tokenstring", "Error": err, "Code": http.StatusInternalServerError})

	}
	claims, err := helper.ToGetClaims(orgtoken)
	if err != nil {
		logger.Logging().Error("Cannot get claims body")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot get claims body", "Error": err, "Code": http.StatusBadRequest})
	}
	userid, err := repository.Dbn.ToGetRedisUserID(claims["uuid"].(string))
	if err != nil {
		return err
	}
	idparam, err := c.ParamsInt("postid")
	if err != nil {
		logger.Logging().Error("Cannot parse query param")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse query param", "Error": err, "Code": http.StatusBadRequest})
	}
	id, err := strconv.Atoi(userid)
	if err != nil {
		return err
	}
	comment.UserId = id
	comment.PostId = idparam
	err = repository.Dbn.ToCreateNewComment(comment)
	if err != nil {
		logger.Logging().Error("Cannot create new comment")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot create new comment", "Error": err, "Code": http.StatusBadRequest})
	}
	logger.Logging().Info("Comment created sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": nil, "Message": "comment created sucessfully", "Error": nil, "Code": http.StatusOK})
}
