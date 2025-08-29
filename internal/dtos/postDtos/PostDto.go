package postDtos

import (
	"Rent-And-Buy-App/internal/dtos/productDtos"
)

type PostDto struct {
	ID       uint                   `json:"id" gorm:"primary_key"`
	SellerId uint                   `json:"SellerId"`
	Price    float32                `json:"Price"`
	Product  productDtos.ProductDto `json:"Product"`
}
