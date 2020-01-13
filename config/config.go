package config

import (
	"os/user"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Common Common `toml:"common"`
	Daily  Daily  `toml:"daily"`
}

type Common struct {
	Root        string `toml:"root"`
	Editor      string `toml:"editor"`
	Finder      string `toml:"finder"`
	FinderOpts  string `toml:"finderOpts"`
	Chooser     string `toml:"chooser"`
	ChooserOpts string `toml:"chooserOpts"`
}

type Daily struct {
	Path         string `toml:"path"`
	TemplatePath string `toml:"template"`
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

	// config.Daily.Path should be a relative path from the config.Common.Root
	// config.Daily.TemplatePath should be a relative path from the config.Common.Root

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

func (c *Config) DailyPath() string {
	return filepath.Join(c.Common.Root, c.Daily.Path)
}

func (c *Config) DailyFilePath(dirPath string, fpath string) string {
	return filepath.Join(
		c.Common.Root,
		c.Daily.Path,
		dirPath,
		fpath,
	)
}

func (c *Config) DailyTemplatePath() string {
	return filepath.Join(c.Common.Root, c.Daily.TemplatePath)
}
