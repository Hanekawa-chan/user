package user

import (
	"github.com/Hanekawa-chan/kanji-user/internal/app"
	"github.com/rs/zerolog"
)

type adapter struct {
	logger *zerolog.Logger
	config *app.Config
	db     app.Database
	auth   app.Auth
}

func NewUser(logger *zerolog.Logger, db app.Database, auth app.Auth, config *app.Config) app.User {
	return &adapter{
		logger: logger,
		db:     db,
		auth:   auth,
		config: config,
	}
}
