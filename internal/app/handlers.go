package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/kanji-team/user/proto/services"
)

func (a *service) CreateUser(ctx context.Context, req *services.CreateUserRequest) (*services.CreateUserResponse, error) {
	id := uuid.New()
	user := &User{
		Id:          id,
		Name:        req.Name,
		Email:       req.Email,
		WordsPerDay: 0,
	}

	err := a.db.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &services.CreateUserResponse{UserId: id.String()}, err
}
