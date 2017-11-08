package scmd

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TOKEN string
}

func LoadToml(confname string) *Config {
	var conf Config

	cur, _ := os.Getwd()
	confpath := path.Join(cur, confname)

	_, err := toml.DecodeFile(confpath, &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
