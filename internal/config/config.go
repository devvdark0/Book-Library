package config

import "os"

type Config struct {
	Env        string
	Addr       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	return &Config{
		Env:        getEnv("Env", "dev"),
		Addr:       getEnv("Address", ":80"),
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
