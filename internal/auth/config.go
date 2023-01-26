package auth

import "time"

type Config struct {
	Address string        `envconfig:"AUTH_ADDRESS"`
	Timeout time.Duration `envconfig:"AUTH_TIMEOUT"`
}
