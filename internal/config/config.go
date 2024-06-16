package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	TgToken string `yaml:"tg_token"`
}

func Load(filePath string) (*Config, error) {
	config := &Config{}
	rawYaml, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYaml, &config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling yaml")
	}

	return config, nil
}
