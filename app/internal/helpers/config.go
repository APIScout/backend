package helpers

import (
	"log"
	"os"

	"backend/app/internal/structs"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config = structs.Config

func LoadConfigs() Config {
	var cfg Config

	// Read config file
	pwd, err := os.Getwd()
	file, err := os.Open(pwd + "/config/local.config.yml")

	// Handle errors
	defer func(file *os.File) {
		err := file.Close()

		if err != nil {
			log.Fatal(err)
		}
	}(file)

	if err != nil {
		log.Fatal(err)
	}

	// Decode yaml file
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	// Decode env variables
	err = envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
