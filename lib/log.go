package lib

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogrus () {
	if err := setFile(Config.Log.File); err != nil {
		log.Fatalf("%v\n", err)
		return
	}
}

func setFile(filePath string) (err error) {
	_, err = os.Stat(filePath)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}

		defer func() {
			err = file.Close()
		}()
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}

	write := bufio.NewWriter(file)
	log.SetOutput(write)
	return
}