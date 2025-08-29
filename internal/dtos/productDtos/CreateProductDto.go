package productDtos

import (
	"Rent-And-Buy-App/internal/dtos/carDtos"
	"Rent-And-Buy-App/internal/dtos/houseDtos"
)

type CreateProductDto struct {
	Car   *carDtos.CreateCarDto     `json:"Car"`
	House *houseDtos.CreateHouseDto `json:"House"`
}
