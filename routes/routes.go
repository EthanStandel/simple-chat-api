package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Build(app fiber.App, router fiber.Router) {
	AuthRoutes(router.Group("/auth"))

	secureRouter := SecureRouter(app, router)
	secureRouter.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
