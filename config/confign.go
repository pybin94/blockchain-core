package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Mongo struct {
		Uri string
		DB  string
	}

	Info struct {
		Version string
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
