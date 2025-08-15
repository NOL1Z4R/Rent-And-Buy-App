package userDtos

type CreateUserDto struct {
	Name     string `json:"Name" binding:"required"`
	Surname  string `json:"Surname" binding:"required"`
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required,min=6"`
	IsSeller *bool  `json:"IsSeller" binding:"required"`
}
