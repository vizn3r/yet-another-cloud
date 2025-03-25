package conf

import (
	"log"
	"os"
)

const (
	MEDIA_DIR = "./media"
)

// Creates required directories and stuff
func Init() error {
	if _, err := os.Stat(MEDIA_DIR); os.IsNotExist(err) {
		log.Println("Creating media directory at", MEDIA_DIR)
		os.Mkdir(MEDIA_DIR, 0755)
	} else {
		return err
	}

	return nil
}
