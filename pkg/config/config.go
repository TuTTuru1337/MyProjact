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

	cfg.DB.DSN = "host=localhost user=postgres password=YOURPASSWORD dbname=postgres port=5432 sslmode=disable"
	cfg.Server.Address = "localhost:8080"

	return &cfg
}
