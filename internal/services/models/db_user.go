package models

type DBUser struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Country     string `db:"country"`
	WordsPerDay int16  `db:"words_per_day"`
}
