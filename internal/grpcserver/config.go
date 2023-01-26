package grpcserver

import "time"

type Config struct {
	Address           string        `envconfig:"GRPC_ADDRESS"`
	MaxConnectionIdle time.Duration `envconfig:"GRPC_MAX_CONNECTION_IDLE"`
	Timeout           time.Duration `envconfig:"GRPC_TIMEOUT"`
	MaxConnectionAge  time.Duration `envconfig:"GRPC_MAX_CONNECTION_AGE"`
	HealthCheckRate   time.Duration `envconfig:"GRPC_HEALTH_CHECK_RATE"`
}
