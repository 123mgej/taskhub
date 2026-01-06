package config

import (
	"os"
)

type Config struct {
	Env  string
	Port string
}

func Load() Config {
	config := Config{
		Env:  getEnv("ENV", "dev"),
		Port: getEnv("PORT", "1000"),
	}
	return config

}

func getEnv(key string, def string) string {
	if v := os.Getenv((key)); v != "" {
		return v
	}
	return def
}
