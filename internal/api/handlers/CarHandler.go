package handlers

import (
	"Rent-And-Buy-App/internal/dtos/carDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/service"
	"Rent-And-Buy-App/pkg/Response"
	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	service *service.CarService
}

func NewCarHandler(service *service.CarService) *CarHandler {
	return &CarHandler{service: service}
}
func (ch *CarHandler) GetCarAll(c *gin.Context) {
	var cars []*entity.Car
	cars, err := ch.service.GetAll()
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	Response.OK(c, gin.H{"cars": cars})
}

func (ch *CarHandler) GetCarById(c *gin.Context) {
	id := c.Param("id")
	car, err := ch.service.GetById(id)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, gin.H{"car": car})
}

func (ch *CarHandler) GetCarByPlate(c *gin.Context) {
	plate := c.Param("plate")
	car, err := ch.service.GetByPlate(plate)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, gin.H{"car": car})
}

func (ch *CarHandler) CreateCar(c *gin.Context) {
	var carDto carDtos.CreateCarDto
	if err := c.ShouldBindJSON(&carDto); err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	car := entity.Car{
		Brand:    carDto.Brand,
		CarModel: carDto.CarModel,
		Year:     carDto.Year,
		Milage:   carDto.Milage,
		Plate:    carDto.Plate,
	}

	car, err := ch.service.CreateCar(&car)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, gin.H{"car": car})
}

func (ch *CarHandler) UpdateCar(c *gin.Context) {
	id := c.Param("id")
	car, err := ch.service.GetById(id)
	if err != nil || car == nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	var carDto *carDtos.UpdateCarDto
	if err := c.ShouldBindJSON(&carDto); err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}

	car, err = ch.service.UpdateCar(car, carDto)
	if err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
		return
	}
	Response.OK(c, gin.H{"car": car})
}

func (ch *CarHandler) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	car, err := ch.service.GetById(id)
	if err != nil || car == nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}

	if err := ch.service.DeleteCar(id); err != nil {
		Response.Error(c, gin.H{"error": err.Error()})
	}

	Response.OK(c, gin.H{"deleted": true})
}
