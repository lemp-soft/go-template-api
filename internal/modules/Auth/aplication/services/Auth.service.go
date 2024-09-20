package services

import (
	AuthDto "go-api-template/internal/modules/Auth/domain/dtos"
	UserService "go-api-template/internal/modules/Users/aplication/service"
	UserDto "go-api-template/internal/modules/Users/domain/dtos"
	UserRepoInterface "go-api-template/internal/modules/Users/domain/repo"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService UserService.UserService
}

// register, login, refresh token, getDatafromToken
func NewAuthService(ur UserRepoInterface.UserRepositoryInterface) AuthService {
	us := UserService.NewUserService(ur)
	return AuthService{
		UserService: us,
	}
}

func (as AuthService) Register(user *UserDto.UserCreateDto) (*UserDto.UserSensitiveDto, error) {
	var userCreateDto UserDto.UserCreateDto = *user
	hash, err := bcrypt.GenerateFromPassword([]byte(userCreateDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userCreateDto.Password = string(hash)
	return as.UserService.Create(user)
}
func (as AuthService) Login(user *AuthDto.LoginDto) (*UserDto.UserSensitiveDto, error) {
	userFound, err := as.UserService.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}
	return userFound, nil
}
