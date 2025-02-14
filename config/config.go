package config

type Config struct {
	PostgresConfig struct {
	}

	ServerConfig struct {
	}

	JWTConfig struct {
		Salt string
	}
}
