package imgs

import (
	"cloud/conf"
	"cloud/log"
	"cloud/util"
	"os"
	"path"

	"github.com/gofiber/fiber/v3"
	"github.com/h2non/filetype"
)

func checkMediaDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("Created 'media' dir")
		}
	}
}

func GET_MediaHandler(c fiber.Ctx) error {
	req_file := c.Params("fname", "")
	if req_file == "" {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}

	file, err := os.ReadFile(path.Join(conf.MEDIA_DIR, req_file))
	if os.IsNotExist(err) {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}

	log.Info("Serving file:", req_file)
	ftype, _ := filetype.Match(file)
	c.Set("Content-Type", ftype.MIME.Value)
	return c.Send(file)
}

func POST_MediaHandler(c fiber.Ctx) error {
	data := c.Body()
	name := util.GenerateHash(32)

	checkMediaDir(conf.MEDIA_DIR)

	log.Info("Writing new file:", path.Join(conf.MEDIA_DIR, name))
	if err := os.WriteFile(path.Join(conf.MEDIA_DIR, name), data, 0644); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendString(name)
}
