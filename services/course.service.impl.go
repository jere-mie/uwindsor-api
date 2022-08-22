package services

import (
	"context"

	"github.com/jere-mie/uwindsor-api/helpers"
	"github.com/jere-mie/uwindsor-api/models"
	"github.com/jmoiron/sqlx"
)

type CourseServiceImpl struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewCourseService(db *sqlx.DB, ctx context.Context) CourseService {
	return &CourseServiceImpl{db: db, ctx: ctx}
}

func (b *CourseServiceImpl) GetCourse(courseCode *string) (*models.Course, error) {
	var courseCodes []string
	b.db.Select(&courseCodes, "SELECT course_code FROM Course")
	closestCode := helpers.ClosestMatch(*courseCode, courseCodes)
	var course = models.Course{}
	err := b.db.Get(&course, "SELECT * FROM Course WHERE course_code = ? LIMIT 1", closestCode)
	return &course, err
}

func (b *CourseServiceImpl) GetCourses() ([]*models.Course, error) {
	var courses []*models.Course
	err := b.db.Select(&courses, "SELECT * FROM Course")
	return courses, err
}
