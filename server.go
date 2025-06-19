package main

import (
	"fmt"

	"github.com/Kimrama/concurrency-mini-chat-terminal/models"
	"github.com/gofiber/fiber/v2"
)

func setupRouter(app *fiber.App) {
	app.Post("/room/:roomname", func(c *fiber.Ctx) error {
		roomName := c.Params("roomname")
		models.Manager.CreateRoom(roomName)
		return c.Status(fiber.StatusOK).SendString("Room " + roomName + "is created")
	})
	app.Post("/room/:roomname/join/:username", func(c *fiber.Ctx) error {
		roomname := c.Params("roomname")
		username := c.Params("username")
		models.Manager.JoinRoom(roomname, username)
		return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Add %s to room[%s]", username, roomname))
	})
	app.Post("/room/:roomname/send", func(c *fiber.Ctx) error {
		type Request struct {
			Message string `json:"msg"`
		}
		roomname := c.Params("roomname")
		requset := new(Request)
		if err := c.BodyParser(requset); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		models.Manager.Broadcast(roomname, requset.Message)
		return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Broadcast message: %s\nTo room: %s", requset.Message, roomname))
	})
}
