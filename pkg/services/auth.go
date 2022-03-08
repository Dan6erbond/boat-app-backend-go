package services

import (
	"golang.org/x/crypto/bcrypt"
	"openwt.com/boat-app-backend/pkg/dto"
	"openwt.com/boat-app-backend/pkg/security"
)

type AuthService interface {
	Login(username, password string) (*dto.LoginResponse, error)
	SignUp(username, password, firstName, lastName string) (*dto.SignUpResponse, error)
}

var _ AuthService = &authService{}

type authService struct {
	usersService UsersService
	jwtUtil      security.JWTUtil
}

func NewAuthService(usersService UsersService, jwtUtil security.JWTUtil) *authService {
	return &authService{usersService: usersService, jwtUtil: jwtUtil}
}

func (s *authService) Login(username, password string) (*dto.LoginResponse, error) {
	var loginResponse *dto.LoginResponse
	user, err := s.usersService.GetUserByUsername(username)
	if err != nil {
		return loginResponse, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return loginResponse, security.ErrInvalidPassword
	}
	accessToken, err := s.jwtUtil.SignAccessToken(security.UserClaims{
		ID:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		return loginResponse, err
	}
	loginResponse = &dto.LoginResponse{
		AccessToken: accessToken,
	}
	return loginResponse, nil
}

func (s *authService) SignUp(username, password, firstName, lastName string) (*dto.SignUpResponse, error) {
	var signUpResponse *dto.SignUpResponse
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return signUpResponse, err
	}
	user, err := s.usersService.CreateUser(&dto.CreateUserDTO{
		Username: username,
		UpdateUserDTO: dto.UpdateUserDTO{
			Password:  string(hashedPassword),
			FirstName: firstName,
			LastName:  lastName,
		},
	})
	if err != nil {
		return signUpResponse, err
	}
	accessToken, err := s.jwtUtil.SignAccessToken(security.UserClaims{
		ID:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		return signUpResponse, err
	}
	signUpResponse = &dto.SignUpResponse{
		LoginResponse: dto.LoginResponse{
			AccessToken: accessToken,
		},
		User: dto.UserDTO{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
	return signUpResponse, nil
}
