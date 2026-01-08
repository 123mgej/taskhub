package config

import (
	"os"
)

type Config struct {
	Env  string
	Port string
	DB_DSN string
}

func Load() Config {
	config := Config{
		Env:  getEnv("ENV", "dev"),
		Port: getEnv("PORT", "1000"),
		DB_DSN: getEnv("DB_DSN","root:root@tcp(127.0.0.1:3306)/taskhub?charset=utf8mb4&parseTime=True&loc=Local"),
	}
	return config

}

func getEnv(key string, def string) string {
	if v := os.Getenv((key)); v != "" {
		return v
	}
	return def
}
