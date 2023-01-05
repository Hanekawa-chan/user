package auth

import (
	"github.com/Hanekawa-chan/kanji-user/internal/app"
	"github.com/rs/zerolog"
	"net/http"
)

type adapter struct {
	logger *zerolog.Logger
	config *app.Config
	client *http.Client
}

func NewAuthClient(logger *zerolog.Logger, config *app.Config) app.Auth {
	client := &http.Client{
		Timeout: config.Auth.Timeout,
	}

	return &adapter{
		logger: logger,
		config: config,
		client: client,
	}
}
