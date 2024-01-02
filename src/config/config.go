package config

import (
	"os"
)

type Config struct {
	MongoURI   string
	APIVersion string
}

func AppConfig() *Config {
	config := &Config{
		MongoURI:   os.Getenv("MONGO_URI"),
		APIVersion: "/api/v1",
	}
	return config
}
