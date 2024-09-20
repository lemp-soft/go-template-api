package service

import (
	UserDto "go-api-template/internal/modules/Users/domain/dtos"
	UserRepoInterface "go-api-template/internal/modules/Users/domain/repo"
)

type UserService struct {
	UserRepository UserRepoInterface.UserRepositoryInterface
}

func NewUserService(ur UserRepoInterface.UserRepositoryInterface) UserService {
	return UserService{
		UserRepository: ur,
	}
}

func (us UserService) Create(user *UserDto.UserCreateDto) (*UserDto.UserSensitiveDto, error) {
	return us.UserRepository.Create(user)
}

func (us UserService) Update(user *UserDto.UserUpdateDto) (*UserDto.UserSensitiveDto, error) {
	return us.UserRepository.Update(user)
}

func (us UserService) Delete(id string) error {
	return us.UserRepository.Delete(id)
}

func (us UserService) FindAll() ([]*UserDto.UserSensitiveDto, error) {
	return us.UserRepository.FindAll()
}

func (us UserService) FindById(id string) (*UserDto.UserSensitiveDto, error) {
	return us.UserRepository.FindById(id)
}

func (us UserService) FindByEmail(email string) (*UserDto.UserSensitiveDto, error) {
	return us.UserRepository.FindByEmail(email)
}
