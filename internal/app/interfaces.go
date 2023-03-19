package app

import (
	"context"
	"user/proto/services"
)

type Service interface {
	CreateUser(ctx context.Context, req *services.CreateUserRequest) (*services.CreateUserResponse, error)
}

type GRPCServer interface {
	ListenAndServe() error
	Shutdown()
}

type Database interface {
	CreateUser(ctx context.Context, req *User) error
}
