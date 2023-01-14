package app

import (
	"context"
	"github.com/Hanekawa-chan/kanji-user/internal/services/models"
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (uuid.UUID, error)
}

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Database interface {
	CreateUser(ctx context.Context, req *models.DBUser) error
}
