package models

import (
	"github.com/google/uuid"
	"github.com/kanji-team/user/internal/app"
)

type User struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	WordsPerDay int16     `db:"words_per_day"`
}

func (u *User) ToDomain() *app.User {
	return &app.User{
		Id:          u.Id,
		Name:        u.Name,
		Email:       u.Email,
		WordsPerDay: u.WordsPerDay,
	}
}

func FromDomain(u *app.User) *User {
	return &User{
		Id:          u.Id,
		Name:        u.Name,
		Email:       u.Email,
		WordsPerDay: u.WordsPerDay,
	}
}
