package config

type Config struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDb       string `env:"POSTGRES_DB"`
	PostgresDbHost   string `env:"POSTGRES_DB_HOST"`
}
