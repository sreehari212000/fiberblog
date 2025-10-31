package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sreehari212000/blog/utils"
)

func CheckAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "token not found in the header")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	id, err := utils.VerifyJwtToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}
	c.Locals("user_id", id)
	return c.Next()
}
