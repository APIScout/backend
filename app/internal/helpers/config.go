package helpers

import (
	"fmt"
	"log"
	"os"

	"backend/app/internal/models"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config = models.Config

// LoadConfigs - parse and store in a struct all the config values needed by the backend.
func LoadConfigs() Config {
	var cfg Config

	// Read config file
	pwd, err := os.Getwd()
	file, err := os.Open(pwd + fmt.Sprintf("/config/%s.config.yml", os.Getenv("GIN_MODE")))

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
