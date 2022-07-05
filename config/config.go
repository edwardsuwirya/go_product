package config

import "os"

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}
type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfigFile() {
	api := os.Getenv("API_URL")
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	c.ApiConfig = ApiConfig{Url: api}
}
func NewConfig() Config {
	cfg := Config{}
	cfg.readConfigFile()
	return cfg
}
