package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Root string `toml:"path"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
