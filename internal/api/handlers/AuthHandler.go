package handlers

import (
	"Rent-And-Buy-App/internal/dtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/service"
	"Rent-And-Buy-App/pkg/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var userDto dtos.CreateUserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		Response.JSON(c, http.StatusUnauthorized, "Validation failed", gin.H{"error": err.Error()})
		return
	}

	user := entity.User{
		Name:     userDto.Name,
		Surname:  userDto.Surname,
		Email:    userDto.Email,
		Password: userDto.Password,
		IsSeller: userDto.IsSeller,
	}

	if err := ah.service.Register(&user); err != nil {
		Response.Error(c, gin.H{"register error": err.Error()})
		return
	}

	Response.OK(c, gin.H{"user": user})
}

func (ah *AuthHandler) Login(c *gin.Context) {

	var userDto dtos.UserLoginDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}

	token, err := ah.service.Login(userDto.Email, userDto.Password)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}
	Response.OK(c, gin.H{"token": token})
}
