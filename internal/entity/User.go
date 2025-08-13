package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"Name" binding:"required"`
	Surname  string `json:"Surname" binding:"required"`
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"-" binding:"required,min=8"`
	IsSeller bool   `json:"IsSeller" binding:"required"`
}
