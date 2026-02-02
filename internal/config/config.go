package config

import "os"

type Config struct {
	AppName string
	Port    string
	DBURL   string
}

func Load() Config {
	return Config{
		AppName: getEnv("APP_NAME", "debt-exempt"),
		Port:    getEnv("APP_PORT", "8080"),
		DBURL:   getEnv("DB_URL", ""),
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
