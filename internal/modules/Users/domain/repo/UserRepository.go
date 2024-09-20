package repo

import UserDto "go-api-template/internal/modules/Users/domain/dtos"

type UserRepositoryInterface interface {
	Create(user *UserDto.UserCreateDto) (*UserDto.UserSensitiveDto, error)
	Update(user *UserDto.UserUpdateDto) (*UserDto.UserSensitiveDto, error)
	Delete(id string) error
	FindAll() ([]*UserDto.UserSensitiveDto, error)
	FindById(id string) (*UserDto.UserSensitiveDto, error)
	FindByEmail(email string) (*UserDto.UserSensitiveDto, error)
}
