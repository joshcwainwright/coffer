package config

import "os"

type Config struct {
	DBPath string
}

func Load() Config {
	return Config{DBPath: env("COFFER_DB_PATH", "./data/coffer.db")}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
