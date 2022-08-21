package services

import (
	"context"

	"github.com/jere-mie/uwindsor-api/helpers"
	"github.com/jere-mie/uwindsor-api/models"
	"github.com/jmoiron/sqlx"
)

type StaffServiceImpl struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewStaffService(db *sqlx.DB, ctx context.Context) StaffService {
	return &StaffServiceImpl{db: db, ctx: ctx}
}

func (b *StaffServiceImpl) GetStaffMember(name *string) (*models.Staff, error) {
	var staffNames []string
	b.db.Select(&staffNames, "SELECT name FROM Staff")
	closestName := helpers.ClosestMatch(*name, staffNames)
	var staffMember = models.Staff{}
	err := b.db.Get(&staffMember, "SELECT * FROM Staff WHERE name = ? LIMIT 1", closestName)
	return &staffMember, err
}

func (b *StaffServiceImpl) GetAllStaff() ([]*models.Staff, error) {
	var allStaff []*models.Staff
	err := b.db.Select(&allStaff, "SELECT * FROM Staff")
	return allStaff, err
}
