package models

import "github.com/google/uuid"

type DBUser struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	Country     string    `db:"country"`
	WordsPerDay int16     `db:"words_per_day"`
}

func (u *DBUser) ToDb(user *CreateUserRequest) {
	u.Country = user.Country
	u.Email = user.Email
}
