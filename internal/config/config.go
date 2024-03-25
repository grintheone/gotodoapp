package config

import "os"

type Config struct {
	Host   string
	Port   string
	DBUser string
	DBPass string
	DBUrl  string
	DBName string
}

func New() *Config {
	return &Config{
		Host:   getEnv("HOST", "localhost"),
		Port:   getEnv("PORT", ":8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASSWORD", ""),
		DBUrl:  getEnv("DB_URL", "127.0.0.1:3306"),
		DBName: getEnv("DB_NAME", "todos"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
