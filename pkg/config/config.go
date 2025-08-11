package config

type Config struct {
	DB struct {
		DSN string
	}
	Server struct {
		Address string
	}
}

func Load() *Config {
	var cfg Config

	cfg.DB.DSN = "postgres://postgres:pass1234@localhost:5432/postgres?sslmode=disable"
	cfg.Server.Address = "localhost:8080"

	return &cfg
}
