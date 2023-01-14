package app

import (
	"github.com/Hanekawa-chan/kanji-user/internal/services/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Logger     *LoggerConfig
	HTTPServer *config.HTTPConfig
	Auth       *config.AuthConfig
	DB         *config.DBConfig
	User       *config.UserConfig
}

type LoggerConfig struct {
	LogLevel string `default:"debug"`
}

func Parse() (*Config, error) {
	cfg := Config{}
	logger := LoggerConfig{}
	db := config.DBConfig{}
	auth := config.AuthConfig{}
	http := config.HTTPConfig{}
	user := config.UserConfig{}
	project := "kanji_auth"

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

	err = envconfig.Process(project, &auth)
	if err != nil {
		log.Err(err).Msg("auth config error")
		return nil, err
	}

	err = envconfig.Process(project, &http)
	if err != nil {
		log.Err(err).Msg("http config error")
		return nil, err
	}

	err = envconfig.Process(project, &user)
	if err != nil {
		log.Err(err).Msg("http config error")
		return nil, err
	}

	cfg.Auth = &auth
	cfg.DB = &db
	cfg.Logger = &logger
	cfg.HTTPServer = &http
	cfg.User = &user

	return &cfg, nil
}
