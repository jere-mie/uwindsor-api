package services

import "github.com/jere-mie/uwindsor-api/models"

type BuildingService interface {
	GetBuilding(*string) (*models.Building, error)
	GetBuildings() ([]*models.Building, error)
}
