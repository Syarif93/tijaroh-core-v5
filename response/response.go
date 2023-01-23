package response

import "github.com/gofiber/fiber/v2"

func Output(c *fiber.Ctx, message string, errors interface{}, item interface{}, items []interface{}) fiber.Map {
	output := fiber.Map{
		"statusCode": c.Response().StatusCode(),
		"message":    message,
		"errors":     errors,
		"data": fiber.Map{
			"item":  item,
			"items": items,
		},
	}

	return output
}
