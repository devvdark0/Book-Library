package config

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	return &Config{
		DBHost:     getEnv("DBHost", "localhost"),
		DBPort:     getEnv("DBPort", "5432"),
		DBUser:     getEnv("DBUser", "user"),
		DBPassword: getEnv("DBPassword", "pass"),
		DBName:     getEnv("DBName", "library"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
