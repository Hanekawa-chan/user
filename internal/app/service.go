package app

import "github.com/rs/zerolog"

type service struct {
	logger *zerolog.Logger
	cfg    *Config
	db     Database
}

func NewService(logger *zerolog.Logger, cfg *Config, db Database) Service {
	return &service{
		logger: logger,
		cfg:    cfg,
		db:     db,
	}
}
