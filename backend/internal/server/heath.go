package server

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func health(c fiber.Ctx) {
	c.JSON(fiber.Map{
		"ok":      true,
		"service": "go-auth",
		"time":    time.Now().UTC(),
	})
}