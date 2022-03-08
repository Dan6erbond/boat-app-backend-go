package repositories

import (
	"errors"

	"gorm.io/gorm"
	"openwt.com/boat-app-backend/pkg/models"
)

var ErrUserNotFound = errors.New("user not found")

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (s *UsersRepository) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UsersRepository) GetUsers() []models.User {
	var users []models.User
	s.db.Find(&users)
	return users
}

func (s *UsersRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return user, ErrUserNotFound
	} else if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UsersRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := s.db.Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, ErrUserNotFound
	} else if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UsersRepository) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

func (s *UsersRepository) DeleteUser(user *models.User) error {
	return s.db.Delete(&user).Error
}

func (s *UsersRepository) DeleteUserByID(id uint) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}
	return s.DeleteUser(&user)
}
