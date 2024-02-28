package helpers

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Elastic struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
	}`yaml:"elastic"`
}


func LoadConfigs() Config {
	var cfg Config

	// Read config file
	pwd, err := os.Getwd()
	file, err := os.Open(pwd + "/local.config.yml")

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
