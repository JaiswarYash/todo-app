package config

import (
	"os"
)

type Config struct {
	MongoURI string
	MongoDB  string
	Port     string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI: os.Getenv("MONGO_URI"),
		MongoDB:  os.Getenv("MONGO_DB"),
		Port:     os.Getenv("PORT"),
	}
}
