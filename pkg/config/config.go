package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Config contains environment specific configuration values.
type Config struct {
	DatabaseDebug bool   `json:"database_debug"`
	DatabaseURL   string `json:"database_url"`
	Environment   string
	Host          string `json:"host"`
	JWTSecret     string `json:"jwt_secret"`
	Port          int    `json:"port"`
}

// LoadConfig returns a Config based on the "ENV" environment variable
func LoadConfig() *Config {
	cfg := Config{}

	switch os.Getenv("ENV") {
	case "production":
		cfg.Environment = "production"
	case "staging":
		cfg.Environment = "staging"
	default:
		cfg.Environment = "development"
	}

	// Heroku environment variables
	if cfg.Environment != "development" {
		cfg.DatabaseURL = os.Getenv("DATABASE_URL")
		cfg.JWTSecret = os.Getenv("JWT_SECRET")

		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
		}
		cfg.Port = port
	}

	file, err := os.Open(fmt.Sprintf("pkg/config/%s.json", cfg.Environment))
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(fileBytes), &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
