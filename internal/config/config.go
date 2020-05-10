package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type List struct {
	Name string `yaml:"name"`
	Id   string `yaml:"id"`
}

type Config struct {
	ApiKey   string `yaml:"api_key"`
	ApiToken string `yaml:"api_token"`
	Lists    []List `yaml:"lists"`
}

func ReadConfig(fileName string) (cfg Config, err error) {
	// open config file
	f, err := os.Open(fileName)
	if err != nil {
		return cfg, fmt.Errorf("could not open config file: %v", err)
	}
	defer f.Close()

	// decode config vars & return as a struct
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	return cfg, err
}
