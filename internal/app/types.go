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
	Id          string
	Name        string
	Email       string
	Country     string
	WordsPerDay int16
}

type CreateUserRequest struct {
	Email   string
	Country string
}

type CreateUserResponse struct {
	UserId string
}
