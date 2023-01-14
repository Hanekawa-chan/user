package app

import "github.com/rs/zerolog"

type service struct {
	logger *zerolog.Logger
	cfg    *Config
	db     Database
	auth   Auth
}

func NewService(logger *zerolog.Logger, cfg *Config, db Database, auth Auth) Service {
	return service{
		logger: logger,
		cfg:    cfg,
		db:     db,
		auth:   auth,
	}
}
