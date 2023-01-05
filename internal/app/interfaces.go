package app

import (
	"context"
	"github.com/Hanekawa-chan/kanji-user/internal/services/models"
)

type Service interface {
}

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Database interface {
}

type User interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error)
}

type Auth interface {
}
