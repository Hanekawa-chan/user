package app

import "github.com/google/uuid"

type User struct {
	Id    uuid.UUID
	Name  string
	Email string
}

type CreateUserRequest struct {
	Email string
	Name  string
}

type CreateUserResponse struct {
	UserId string
}
