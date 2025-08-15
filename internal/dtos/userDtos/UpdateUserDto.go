package userDtos

type UpdateUserDto struct {
	Name     string `json:"Name" binding:"required"`
	Surname  string `json:"Surname" binding:"required"`
	Email    string `json:"Email" binding:"required,email"`
	IsSeller *bool  `json:"IsSeller"`
}
