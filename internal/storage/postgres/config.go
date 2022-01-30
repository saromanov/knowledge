package postgres

// Config defines configuration for postgres connection
type Config struct {
	Username string `env:"POSTGRES_USERNAME"`
	Password string `env:"POSTGRES_PASSWORD"`
	DB string `env:"POSTGRES_DB"`
}