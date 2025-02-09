package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Port               string `json:"PORT"`
	DatabaseURL        string `json:"DATABASE_URL"`
	JWTsignatureSecret string `json:"JWT_SIGNATURE_SECRET"`
}

var Config Configuration

func LoadConfig() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		log.Fatalf("error decoding config JSON: %v", err)
		return err
	}

	// TODO: override with .env parameters
	return nil
}
