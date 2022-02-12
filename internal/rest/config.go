package rest

import "time"

// Config defines configuration for rest
type Config struct {
	Address         string        `env:"REST_ADDRESS"`
	ShutdownTimeout time.Duration `env:"REST_SHUTDOWN_TIMEOUT,default=30s"`
}
