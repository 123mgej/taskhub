package config

import (
	"os"
	"strconv"
)

type Config struct {
	Env  string
	Port string
	DB_DSN string

	JWTSecret string
	JWTExpireMinutes int
}

func Load() Config {
	config := Config{
		Env:  getEnv("ENV", "dev"),
		Port: getEnv("PORT", "1000"),
		DB_DSN: getEnv("DB_DSN","root:root@tcp(127.0.0.1:3306)/taskhub?charset=utf8mb4&parseTime=True&loc=Local"),
		JWTSecret:        getEnv("JWT_SECRET", "dev-secret-please-change"),
		JWTExpireMinutes: getEnvInt("JWT_EXPIRE_MINUTES", 120),
	}
	return config

}

func getEnv(key string, def string) string {
	if v := os.Getenv((key)); v != "" {
		return v
	}
	return def
}



func getEnvInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil || n <= 0 {
		return def
	}
	return n
}
