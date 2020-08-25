package main

import "github.com/BurntSushi/toml"

type LogFile struct {
	Path string `toml:"path"`
}

type Config struct {
	Port int                `toml:"port"`
	Logs map[string]LogFile `toml:"log"`
}

func readConfig(path string) (Config, error) {
	var c Config

	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
