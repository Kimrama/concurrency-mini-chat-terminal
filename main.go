package main

import "github.com/gofiber/fiber/v2"

func main() {
	go interactive()

	app := fiber.New()
	setupRouter(app)

	app.Listen(":8080")
}
