package config

import "github.com/kelseyhightower/envconfig"

type AppConfig struct {
	DBHost           string `envconfig:"DB_HOST" required:"true"`
	DBPort           string `envconfig:"DB_PORT" required:"true"`
	DBName           string `envconfig:"DB_NAME" required:"true"`
	DBUser           string `envconfig:"DB_USER" required:"true"`
	DBPassword       string `envconfig:"DB_PASSWORD" required:"true"`
	DBMigrationsPath string `envconfig:"DB_MIGRATIONS_PATH" required:"true"`
	DBSSLMode        string `envconfig:"PGSQLMODE" default:"disable"`
}

// GetAppConfigFromEnv using environment variables
func GetAppConfigFromEnv() (AppConfig, error) {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
