package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"user/internal/app"
	"user/internal/database"
	"user/internal/grpcserver"
)

type Config struct {
	Logger     *LoggerConfig
	GRPCServer *grpcserver.Config
	DB         *database.Config
	User       *app.Config
}

type LoggerConfig struct {
	LogLevel string `default:"debug"`
}

func Parse() (*Config, error) {
	cfg := Config{}
	logger := LoggerConfig{}
	db := database.Config{}
	grpc := grpcserver.Config{}
	user := app.Config{}
	project := "KANJI_USER"

	err := envconfig.Process(project, &logger)
	if err != nil {
		log.Err(err).Msg("logger config error")
		return nil, err
	}

	err = envconfig.Process(project, &db)
	if err != nil {
		log.Err(err).Msg("db config error")
		return nil, err
	}

	err = envconfig.Process(project, &grpc)
	if err != nil {
		log.Err(err).Msg("http config error")
		return nil, err
	}

	err = envconfig.Process(project, &user)
	if err != nil {
		log.Err(err).Msg("http config error")
		return nil, err
	}

	cfg.DB = &db
	cfg.Logger = &logger
	cfg.GRPCServer = &grpc
	cfg.User = &user

	return &cfg, nil
}
