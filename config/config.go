package config

import (
	"github.com/BurntSushi/toml"
)

func LoadConfig(filename string, c interface{}) error {
	_, err := toml.DecodeFile(filename, c)
	return err
}
