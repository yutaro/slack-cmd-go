package scmd

import (
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TOKEN string
}

func LoadToml(confname string) *Config {
	var conf Config
	_, cur, _, _ := runtime.Caller(1)
	confpath := path.Join(path.Dir(cur), confname)

	_, err := toml.DecodeFile(confpath, &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
