package main

import (
	"cloud/conf"
	"cloud/imgs"
	"cloud/log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	log.Init(true)
	conf.Init()
	app := fiber.New(fiber.Config{})

	app.All("/", func(c fiber.Ctx) error {
		return c.SendString("im fine :)")
	})

	app.Get("/m/:fname", imgs.GET_MediaHandler)
	app.Post("/m", imgs.POST_MediaHandler)

	app.Listen(":8080", fiber.ListenConfig{})
}
