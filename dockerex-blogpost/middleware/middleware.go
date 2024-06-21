package middleware

import (
	"blogpost/helper"
	"blogpost/repository"
	"fmt"
	"strings"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AutherizationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Middleware Trigred")
		tokenHeader := c.Get(fiber.HeaderAuthorization)
		if tokenHeader == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		uuidclaims, err := helper.ToGetClaims(tokenHeader[7:])

		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT token"})
		}
		uuid := uuidclaims["uuid"]
		role, err := repository.Dbn.ToGetRedisCache(uuid.(string))
		if err != nil {
			return nil
		}

		path := c.Path()

		fmt.Println("path", path)
		fmt.Println("role---->", role)

		if strings.Contains(path, "/admin/") && role == "user" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized user"})
		}
		c.Set(fiber.HeaderAuthorization, tokenHeader)
		c.Context().SetUserValue("uuid", tokenHeader[7:])
		return c.Next()
	}

}
