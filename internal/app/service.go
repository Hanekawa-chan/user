package app

import "github.com/rs/zerolog"

type service struct {
	logger *zerolog.Logger
	cfg    *Config
	user   User
}

func NewService(logger *zerolog.Logger, cfg *Config, user User) Service {
	return service{
		logger: logger,
		cfg:    cfg,
		user:   user,
	}
}
