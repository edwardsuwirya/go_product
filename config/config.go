package config

import "os"

type ApiConfig struct {
	Url string
}
type Config struct {
	ApiConfig
}

func (c *Config) readConfigFile() {
	api := os.Getenv("API_URL")
	c.ApiConfig = ApiConfig{Url: api}
}
func NewConfig() Config {
	cfg := Config{}
	cfg.readConfigFile()
	return cfg
}
