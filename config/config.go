package config

import (
	"fmt"
	"log"
	"os"
	"shop-service/pkg/logger"

	"gopkg.in/yaml.v3"
)

type JWTConfig struct {
	Secret          string `yaml:"secret"`
	ExpirationHours int    `yaml:"expiration_hours"`
}

type Config struct {
	PostgresConfig `yaml:"postgres"`
	JWTConfig      `yaml:"jwt"`
	ServerConfig   struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	logger.Log.Debug().Str("env", env).Send()

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
