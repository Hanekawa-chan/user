package models

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
