package configuration

import "github.com/OE-OverEngineer/over-review-backend/utils/env"

// Config is a struct that contains configuration variables
type Config struct {
	Port        string
	SecretKey   string
	Environment string
	Database    *Database
}

// Database is a struct that contains DB's configuration variables
type Database struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	port := env.MustGet("PORT")
	// set default PORT if missing
	if port == "" {
		port = "3000"
	}

	secretKey := env.MustGet("SECRET_KEY")
	environment := env.MustGet("Env")
	if environment != "Production" {
		secretKey = "local"
	}
	return &Config{
		Port:        port,
		SecretKey:   secretKey,
		Environment: environment,
		Database: &Database{
			Host:     env.MustGet("DATABASE_HOST"),
			Port:     env.MustGet("DATABASE_PORT"),
			User:     env.MustGet("DATABASE_USER"),
			Name:     env.MustGet("DATABASE_NAME"),
			Password: env.MustGet("DATABASE_PASS"),
		},
	}, nil
}
