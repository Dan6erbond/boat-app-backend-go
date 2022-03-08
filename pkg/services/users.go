package services

import (
	dtoMapper "github.com/dranikpg/dto-mapper"
	"openwt.com/boat-app-backend/pkg/dto"
	"openwt.com/boat-app-backend/pkg/models"
	"openwt.com/boat-app-backend/pkg/repositories"
)

type UsersService interface {
	CreateUser(createDTO *dto.CreateUserDTO) (*models.User, error)
	GetUsers() []models.User
	GetUserByID(id uint) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	UpdateUser(id uint, updateDTO *dto.UpdateUserDTO) (*models.User, error)
	DeleteUserByID(id uint) error
}

var _ UsersService = &usersService{}

type usersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) *usersService {
	return &usersService{usersRepository: usersRepository}
}

func (s *usersService) CreateUser(createDTO *dto.CreateUserDTO) (*models.User, error) {
	var user models.User
	dtoMapper.Map(&user, createDTO)
	err := s.usersRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (s *usersService) GetUsers() []models.User {
	return s.usersRepository.GetUsers()
}

func (s *usersService) GetUserByID(id uint) (models.User, error) {
	return s.usersRepository.GetUserByID(id)
}

func (s *usersService) GetUserByUsername(username string) (models.User, error) {
	return s.usersRepository.GetUserByUsername(username)
}

func (s *usersService) UpdateUser(id uint, updateDTO *dto.UpdateUserDTO) (*models.User, error) {
	user, err := s.usersRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	dtoMapper.Map(&user, updateDTO)
	err = s.usersRepository.UpdateUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (s *usersService) DeleteUserByID(id uint) error {
	return s.usersRepository.DeleteUserByID(id)
}
