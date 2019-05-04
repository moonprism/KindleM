package lib

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogrus () {
	if err := setLogFile(); err != nil {
		log.Fatalf("%v\n", err)
		return
	}
}

func setLogFile() (err error) {
	_, err = os.Stat(Config.Log.File)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(Config.Log.File)
		if err != nil {
			return err
		}

		defer func() {
			err = file.Close()
		}()
	}

	file, err := os.OpenFile(Config.Log.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}

	write := bufio.NewWriter(file)
	log.SetOutput(write)
	return
}