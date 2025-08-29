package productDtos

import (
	"Rent-And-Buy-App/internal/dtos/carDtos"
	"Rent-And-Buy-App/internal/dtos/houseDtos"
)

type ProductDto struct {
	ID    uint                `json:"ID" gorm:"primary_key"`
	Car   *carDtos.CarDto     `json:"Car"`
	House *houseDtos.HouseDto `json:"House"`
}
