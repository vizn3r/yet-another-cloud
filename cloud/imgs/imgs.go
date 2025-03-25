package imgs

import (
	"cloud/conf"
	"cloud/util"
	"log"
	"os"
	"path"

	"github.com/gofiber/fiber/v3"
	"github.com/h2non/filetype"
)

func GET_MediaHandler(c fiber.Ctx) error {
	req_file := c.Params("fname", "")
	if req_file == "" {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}

	file, err := os.ReadFile(path.Join(conf.MEDIA_DIR, req_file))
	if os.IsNotExist(err) {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}

	ftype, _ := filetype.Match(file)
	c.Set("Content-Type", ftype.MIME.Value)
	return c.Send(file)
}

func POST_MediaHandler(c fiber.Ctx) error {
	data := c.Body()
	name := util.GenerateHash(32)

	log.Println("Writing new file:", name)
	os.WriteFile(path.Join(conf.MEDIA_DIR, name), data, 0644)
	return c.SendString(name)
}
