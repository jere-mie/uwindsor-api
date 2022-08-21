package services

import "github.com/jere-mie/uwindsor-api/models"

type CourseService interface {
	GetCourse(*string) (*models.Course, error)
	GetCourses() ([]*models.Course, error)
}
