package scmd

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TOKEN string
}

func LoadToml(path string) *Config {
	var conf Config
	_, err := toml.DecodeFile(path, &conf)
	if err == nil {
		return &conf
	}

	basepath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fullpath := filepath.Join(basepath, path)

	_, err = toml.DecodeFile(fullpath, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
