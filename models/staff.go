package models

type Staff struct {
	Name       string `db:"name" json:"name"`
	Title      string `db:"title" json:"title"`
	Phone      string `db:"phone" json:"phone"`
	Email      string `db:"email" json:"email"`
	Location   string `db:"location" json:"location"`
	Department string `db:"department" json:"department"`
}
