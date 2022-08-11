package main

import (
	"log"
	"os"
	"simple-chat-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Build(*app, app.Group("/api"))
	log.Fatal(app.Listen(os.Getenv("SCA_PORT")))
}
