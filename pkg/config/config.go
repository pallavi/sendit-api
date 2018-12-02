package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config contains environment specific configuration values.
type Config struct {
	DatabaseHost     string `json:"database_host"`
	DatabaseName     string `json:"database_name"`
	DatabasePassword string `json:"database_password"`
	DatabasePort     int    `json:"database_port"`
	DatabaseUser     string `json:"database_user"`
	Environment      string
	Host             string `json:"host"`
	JWTSecret        string `json:"jwt_secret"`
	Port             int    `json:"port"`
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
