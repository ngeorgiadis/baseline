package config

// AppConfig ...
type AppConfig struct {
	Port     int    `toml:"port"`
	RootPath string `toml:"root_path"`
}

// Config is the main config type
type Config struct {
	App AppConfig `toml:"app"`
}
