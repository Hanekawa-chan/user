package models

import (
	"github.com/google/uuid"
	"user/internal/app"
)

type User struct {
	Id    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	Email string    `db:"email"`
}

func (u *User) ToDomain() *app.User {
	return &app.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}
}

func FromDomain(u *app.User) *User {
	return &User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}
}
