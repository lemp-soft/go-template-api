package repo

import (
	DbConfig "go-api-template/internal/core/infra/config"
	UserDto "go-api-template/internal/modules/Users/domain/dtos"
	UserEntity "go-api-template/internal/modules/Users/domain/entitys"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepositoryImpl {
	return UserRepositoryImpl{}
}

func (ur UserRepositoryImpl) Create(user *UserDto.UserCreateDto) (*UserDto.UserSensitiveDto, error) {
	newUser := UserEntity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	UserCreate := DbConfig.Db.Create(&newUser)
	if UserCreate.Error != nil {
		return nil, UserCreate.Error
	}
	return &UserDto.UserSensitiveDto{
		ID:       newUser.ID.String(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	}, nil
}
func (ur UserRepositoryImpl) Update(user *UserDto.UserUpdateDto) (*UserDto.UserSensitiveDto, error) {
	newUser := UserEntity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	UserUpdate := DbConfig.Db.Save(&newUser)
	if UserUpdate.Error != nil {
		return nil, UserUpdate.Error
	}
	return &UserDto.UserSensitiveDto{
		ID:       newUser.ID.String(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	}, nil
}
func (ur UserRepositoryImpl) Delete(id string) error {
	UserDelete := DbConfig.Db.Delete(&UserEntity.User{}, id)
	if UserDelete.Error != nil {
		return UserDelete.Error
	}
	return nil
}
func (ur UserRepositoryImpl) FindAll() ([]*UserDto.UserSensitiveDto, error) {
	var users []*UserEntity.User
	UserFindAll := DbConfig.Db.Find(&users)
	if UserFindAll.Error != nil {
		return nil, UserFindAll.Error
	}
	var usersDto []*UserDto.UserSensitiveDto
	for _, user := range users {
		usersDto = append(usersDto, &UserDto.UserSensitiveDto{
			ID:       user.ID.String(),
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}
	return usersDto, nil
}

func (ur UserRepositoryImpl) FindById(id string) (*UserDto.UserSensitiveDto, error) {
	var user UserEntity.User
	UserFindById := DbConfig.Db.Where("id = ?", id).First(&user)
	if UserFindById.Error != nil {
		return nil, UserFindById.Error
	}
	return &UserDto.UserSensitiveDto{
		ID:       user.ID.String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (ur UserRepositoryImpl) FindByEmail(email string) (*UserDto.UserSensitiveDto, error) {
	var user UserEntity.User
	UserFindByEmail := DbConfig.Db.Where("email = ?", email).First(&user)
	if UserFindByEmail.Error != nil {
		return nil, UserFindByEmail.Error
	}
	return &UserDto.UserSensitiveDto{
		ID:       user.ID.String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
