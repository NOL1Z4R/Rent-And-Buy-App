package handlers

import (
	"Rent-And-Buy-App/internal/dtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/service"
	"Rent-And-Buy-App/pkg/Response"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (uh *UserHandler) GetAll(c *gin.Context) {
	var users []entity.User
	var err error
	users, err = uh.service.GetAll()

	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}

	Response.OK(c, gin.H{"users": users})
}

func (uh *UserHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.service.GetByID(id)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	Response.OK(c, gin.H{"user": user})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var updateDto dtos.UpdateUserDto
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}

	id := c.Param("id")
	user, err := uh.service.GetByID(id)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}

	updatedUser, err := uh.service.UpdateUser(user, updateDto)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}
	Response.OK(c, gin.H{"user": updatedUser})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if strings.Trim(id, " ") == "" {
		Response.Error(c, gin.H{"error": "id is required"})
	}
	err := uh.service.DeleteUser(id)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}
	Response.OK(c, gin.H{"data": true})
}
