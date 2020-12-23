package config

import "github.com/BurntSushi/toml"

// New ...
func New(file string) (*Config, error) {

	var c *Config

	_, err := toml.DecodeFile(file, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
