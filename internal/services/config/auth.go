package config

import "time"

type AuthConfig struct {
	Address string        `envconfig:"USER_ADDRESS"`
	Timeout time.Duration `envconfig:"USER_TIMEOUT"`
}
