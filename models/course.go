package models

type Course struct {
	Course_Code string `db:"course_code" json:"course_code"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}
