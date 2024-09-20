package dtos

// type User struct {
// 	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
// 	Name     string    `json:"name" gorm:"not null"`
// 	Email    string    `json:"email" gorm:"unique;not null"`
// 	Password string    `json:"password" gorm:"not null"`
// }

type UserCreateDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type UserUpdateDto struct {
	Name     string `json:"name" validate:"string"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"string"`
}
type UserNotSensitiveDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserSensitiveDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
