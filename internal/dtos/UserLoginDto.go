package dtos

type UserLoginDto struct {
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required,min=6"`
}
