package postgres

// Config defines configuration for postgres connection
type Config struct {
	Host           string `env:"POSTGRES_HOST"`
	Port           int    `env:"POSTGRES_PORT,default=5432"`
	Username       string `env:"POSTGRES_USERNAME"`
	Password       string `env:"POSTGRES_PASSWORD"`
	DB             string `env:"POSTGRES_DB"`
	ConnectRetries int    `env:POSTGERS_CONNECT_RETRIES,default=3`
}
