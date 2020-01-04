package config

import (
	"os/user"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Common Common `toml:"common"`
}

type Common struct {
	Root   string `toml:"root"`
	Editor string `toml:"editor"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	var err error

	config.Common.Root, err = expandPath(config.Common.Root)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func expandPath(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return filepath.Abs(path)
	}

	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return strings.Replace(path, "~", user.HomeDir, 1), nil
}
