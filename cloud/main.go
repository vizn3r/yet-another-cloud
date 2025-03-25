package main

import (
	"cloud/conf"
	"cloud/imgs"

	"github.com/gofiber/fiber/v3"
)

func main() {
	conf.Init()
	app := fiber.New(fiber.Config{})

	app.All("/", func(c fiber.Ctx) error {
		return c.SendString("im fine :)")
	})

	app.Get("/m/:fname", imgs.GET_MediaHandler)
	app.Post("/m", imgs.POST_MediaHandler)

	app.Listen(":8080", fiber.ListenConfig{})
}
