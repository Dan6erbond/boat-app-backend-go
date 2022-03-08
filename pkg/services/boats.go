package services

import (
	dtoMapper "github.com/dranikpg/dto-mapper"
	"openwt.com/boat-app-backend/pkg/dto"
	"openwt.com/boat-app-backend/pkg/models"
	"openwt.com/boat-app-backend/pkg/repositories"
)

type BoatsService interface {
	GetBoats() []models.Boat
	GetBoatByID(id uint) (models.Boat, error)
	CreateBoat(createDTO *dto.CreateUpdateBoatDTO) (*models.Boat, error)
	UpdateBoat(id uint, updateDTO *dto.CreateUpdateBoatDTO) (*models.Boat, error)
	DeleteBoatByID(id uint) error
}

var _ BoatsService = &boatsService{}

type boatsService struct {
	boatsRepository *repositories.BoatsRepository
}

func NewBoatsService(boatsRepository *repositories.BoatsRepository) *boatsService {
	return &boatsService{boatsRepository: boatsRepository}
}

func (s *boatsService) CreateBoat(createDTO *dto.CreateUpdateBoatDTO) (*models.Boat, error) {
	var boat models.Boat
	dtoMapper.Map(&boat, createDTO)
	err := s.boatsRepository.CreateBoat(&boat)
	if err != nil {
		return nil, err
	}
	return &boat, err
}

func (s *boatsService) GetBoats() []models.Boat {
	return s.boatsRepository.GetBoats()
}

func (s *boatsService) GetBoatByID(id uint) (models.Boat, error) {
	return s.boatsRepository.GetBoatByID(id)
}

func (s *boatsService) UpdateBoat(id uint, updateDTO *dto.CreateUpdateBoatDTO) (*models.Boat, error) {
	boat, err := s.boatsRepository.GetBoatByID(id)
	if err != nil {
		return nil, err
	}
	dtoMapper.Map(&boat, updateDTO)
	err = s.boatsRepository.UpdateBoat(&boat)
	if err != nil {
		return nil, err
	}
	return &boat, err
}

func (s *boatsService) DeleteBoatByID(id uint) error {
	return s.boatsRepository.DeleteBoatByID(id)
}
