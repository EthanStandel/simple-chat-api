package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Build(router fiber.Router) {
	Auth(router.Group("/auth"))
}
