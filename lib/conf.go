package lib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var Config struct {
	Storage struct {
		Host string `yaml:"host"`
		Port string	`yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"storage"`
	Log struct {
		File string `yaml:"file"`
		Level string `yaml:"level"`
	} `yaml:"log"`
}

const RuntimeDir = "./runtime"

// init parse config should be run in the beginning
func init() {
	content, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatalf("read config : %v", err)
	}

	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		log.Fatalf("unmarshal : %v", err)
	}

	// create runtime dir
	_, err = os.Stat(RuntimeDir)

	if os.IsNotExist(err) {
		err = os.Mkdir(RuntimeDir, os.ModePerm)

		if err != nil {
			log.Fatalf("mkdir runtime dir failed : %v", err)
		}
	}
}
