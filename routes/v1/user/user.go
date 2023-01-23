package user

import (
	"core-v5/response"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(user fiber.Router) fiber.Router {
	user.Get("/user/list", func(c *fiber.Ctx) error {
		return c.JSON(response.Output(
			c, "User List", nil, nil, []interface{}{
				fiber.Map{"id": 1, "name": "Moh. Sarifudin"},
			},
		))
	})

	return user
}
