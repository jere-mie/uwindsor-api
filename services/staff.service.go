package services

import "github.com/jere-mie/uwindsor-api/models"

type StaffService interface {
	GetStaffMember(*string) (*models.Staff, error)
	GetAllStaff() ([]*models.Staff, error)
}
