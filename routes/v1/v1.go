package v1

import (
	"core-v5/routes/v1/user"

	"github.com/gofiber/fiber/v2"
)

func ApiV1(v1 fiber.Router) fiber.Router {
	user.UserRoutes(v1)

	return v1
}
