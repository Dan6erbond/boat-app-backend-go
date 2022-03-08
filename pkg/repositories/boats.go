package repositories

import (
	"fmt"

	"gorm.io/gorm"
	"openwt.com/boat-app-backend/pkg/models"
)

var ErrBoatNotFound error = fmt.Errorf("boat not found")

type BoatsRepository struct {
	db *gorm.DB
}

func NewBoatsRepository(db *gorm.DB) *BoatsRepository {
	return &BoatsRepository{db: db}
}

func (s *BoatsRepository) CreateBoat(boat *models.Boat) error {
	return s.db.Create(boat).Error
}

func (s *BoatsRepository) GetBoats() []models.Boat {
	var boats []models.Boat
	s.db.Find(&boats)
	return boats
}

func (s *BoatsRepository) GetBoatByID(id uint) (models.Boat, error) {
	var boat models.Boat
	err := s.db.First(&boat, id).Error
	if err == gorm.ErrRecordNotFound {
		return boat, ErrBoatNotFound
	} else if err != nil {
		return boat, err
	}
	return boat, nil
}

func (s *BoatsRepository) UpdateBoat(boat *models.Boat) error {
	return s.db.Save(boat).Error
}

func (s *BoatsRepository) DeleteBoat(boat *models.Boat) error {
	return s.db.Delete(&boat).Error
}

func (s *BoatsRepository) DeleteBoatByID(id uint) error {
	boat, err := s.GetBoatByID(id)
	if err != nil {
		return err
	}
	return s.DeleteBoat(&boat)
}
