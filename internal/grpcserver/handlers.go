package grpcserver

import (
	"context"
	"github.com/kanji-team/user/proto/services"
)

func (a *adapter) CreateUser(ctx context.Context, request *services.CreateUserRequest) (*services.CreateUserResponse, error) {
	return a.service.CreateUser(ctx, request)
}

func (a *adapter) GetUserInfo(ctx context.Context, request *services.UserInfoRequest) (*services.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *adapter) ChangeUsername(ctx context.Context, request *services.ChangeNameRequest) (*services.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (a *adapter) Check(ctx context.Context, request *services.HealthCheckRequest) (*services.HealthCheckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *adapter) Watch(request *services.HealthCheckRequest, server services.Health_WatchServer) error {
	//TODO implement me
	panic("implement me")
}
