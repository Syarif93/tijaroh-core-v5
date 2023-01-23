package main

import (
	"os"

	"core-v5/process"
	"core-v5/response"
	v1 "core-v5/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	os.Setenv("TZ", process.Env().TimeZone)
}

func main() {
	file := process.LoggerFile()
	defer file.Close()

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(response.Output(
			c, "Server Online", nil, nil, nil,
		))
	})

	v1Group := app.Group("/api/v1")
	v1.ApiV1(v1Group)

	app.Listen(":" + process.Env().AppPort)
}
