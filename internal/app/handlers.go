package app

import (
	"context"
	"github.com/Hanekawa-chan/kanji-user/internal/services/models"
	"github.com/google/uuid"
)

func (a *service) CreateUser(ctx context.Context, req *models.CreateUserRequest) (uuid.UUID, error) {
	id := uuid.New()
	var user models.DBUser
	user.ToDb(req)
	user.Id = id

	err := a.db.CreateUser(ctx, &user)
	if err != nil {
		return [16]byte{}, err
	}

	return id, err
}
