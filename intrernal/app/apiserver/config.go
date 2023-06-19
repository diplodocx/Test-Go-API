package apiserver

import "github.com/BurntSushi/toml"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig(filepath string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(filepath, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
