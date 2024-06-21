package handler

import (
	"blogpost/helper"
	"blogpost/logger"
	"blogpost/models"
	"blogpost/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *fiber.Ctx) error {
	userlogin := new(models.Logincredentials)
	err := c.BodyParser(userlogin)
	if err != nil {
		logger.Logging().Error("Cannot parse request body")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Data": nil, "Message": "Cannot parse request body", "Error": err, "Code": http.StatusInternalServerError})
	}
	data, err := repository.Dbn.ToLoginUser(userlogin)
	if err != nil {
		logger.Logging().Error("Cannot invalid user credentials")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot invalid user credentials", "Error": err, "Code": http.StatusBadRequest})
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(userlogin.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "incorrect password", "Error": err, "Code": http.StatusBadRequest})

	}
	uniqueid := helper.Uuidgen()
	uuid, err := helper.CreateTokenWithoutClaim(uniqueid)
	if err != nil {
		logger.Logging().Error("Not able to create token withoutclaims")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Not able to create token withoutclaims", "Error": err, "Code": http.StatusBadRequest})
	}

	if err := repository.Dbn.ToSetRedisCache(data, uniqueid); err != nil {
		logger.Logging().Error("Cannot invalid user credentials")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": nil, "Message": "Cannot set redis cache", "Error": err, "Code": http.StatusBadRequest})
	}

	logger.Logging().Info("User login sucessfully")
	return c.Status(http.StatusOK).JSON(fiber.Map{"Data": uuid, "Message": "User login sucessfull", "Error": nil, "Code": http.StatusOK})
}
