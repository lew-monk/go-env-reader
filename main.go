package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadConfig[T any](config *T) (*T, error) {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, skipping...")
	}

	cfg, err := env.ParseAs[T]()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse config: %w", err)
	}

	return &cfg, nil
}
