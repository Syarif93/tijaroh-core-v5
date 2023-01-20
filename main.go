package main

import (
	"os"

	"core-v5/process"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	os.Setenv("TZ", process.Env().TimeZone)
}

func main() {
	file := process.LoggerFile()

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	app.Listen(":" + process.Env().AppPort)
}
