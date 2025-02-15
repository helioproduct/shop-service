package config

type Config struct {
	PostgresConfig struct {
	}

	ServerConfig struct {
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
