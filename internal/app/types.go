package app

import "github.com/google/uuid"

type Word struct {
	Id       int16
	Text     string
	Hiragana string
	Kanjis   []Kanji
}

type Kanji struct {
	Id       uuid.UUID
	Kanji    string
	Hiragana string
}

type User struct {
	Id          uuid.UUID
	Name        string
	Email       string
	WordsPerDay int16
}

type CreateUserRequest struct {
	Email string
	Name  string
}

type CreateUserResponse struct {
	UserId string
}
