package handlers

import (
	"Rent-And-Buy-App/internal/dtos/houseDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/service"
	"Rent-And-Buy-App/pkg/Response"
	"github.com/gin-gonic/gin"
)

type HouseHandler struct {
	service service.HouseService
}

func NewHouseHandler(service service.HouseService) *HouseHandler {
	return &HouseHandler{service: service}
}

func (hh *HouseHandler) CreateHouse(c *gin.Context) {
	var createDto houseDtos.CreateHouseDto
	if err := c.ShouldBindJSON(&createDto); err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	house := entity.House{
		Adress:      createDto.Adress,
		Year:        createDto.Year,
		SquareMeter: createDto.SquareMeter,
	}

	newHouse, err := hh.service.Create(&house)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, newHouse)
}

func (hh *HouseHandler) GetHouse(c *gin.Context) {
	id := c.Param("id")
	house, err := hh.service.GetById(id)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, house)
}

func (hh *HouseHandler) GetHouseAll(c *gin.Context) {
	house, err := hh.service.GetAll()
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, house)
}

func (hh *HouseHandler) UpdateHouse(c *gin.Context) {
	id := c.Param("id")
	house, err := hh.service.GetById(id)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	var updateDto houseDtos.UpdateHouseDto
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	newHouse, err := hh.service.Update(house, &updateDto)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, newHouse)
}

func (hh *HouseHandler) DeleteHouse(c *gin.Context) {
	id := c.Param("id")
	house, err := hh.service.GetById(id)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	err = hh.service.Delete(house)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, gin.H{"data": "success"})
}
