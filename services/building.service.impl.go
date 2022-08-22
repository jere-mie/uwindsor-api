package services

import (
	"context"

	"github.com/jere-mie/uwindsor-api/helpers"
	"github.com/jere-mie/uwindsor-api/models"
	"github.com/jmoiron/sqlx"
)

type BuildingServiceImpl struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewBuildingService(db *sqlx.DB, ctx context.Context) BuildingService {
	return &BuildingServiceImpl{db: db, ctx: ctx}
}

func (b *BuildingServiceImpl) GetBuilding(name *string) (*models.Building, error) {
	var buildingNames []string
	b.db.Select(&buildingNames, "SELECT building_name FROM Building")
	closestName := helpers.ClosestMatch(*name, buildingNames)
	var building = models.Building{}
	err := b.db.Get(&building, "SELECT * FROM Building WHERE building_name = ? LIMIT 1", closestName)
	return &building, err
}

func (b *BuildingServiceImpl) GetBuildings() ([]*models.Building, error) {
	var buildings []*models.Building
	err := b.db.Select(&buildings, "SELECT * FROM Building")
	return buildings, err
}
