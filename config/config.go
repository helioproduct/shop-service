package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PostgresConfig `yaml:"postgres"`
	ServerConfig   struct {
	}
	JWTConfig struct {
		Salt string
	}
	Tables struct {
		Users     string
		Products  string
		Purchases string
		Transfers string
	}
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	configFile := fmt.Sprintf("./config/%s.yaml", env)

	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("error opening config file %s: %w", configFile, err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("error decoding config file %s: %w", configFile, err)
	}

	log.Printf("Loaded configuration from %s", configFile)
	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
