package productDtos

import (
	"Rent-And-Buy-App/internal/dtos/carDtos"
	"Rent-And-Buy-App/internal/dtos/houseDtos"
)

type UpdateProductDto struct {
	Car   *carDtos.UpdateCarDto     `json:"Car"`
	House *houseDtos.UpdateHouseDto `json:"House"`
}
