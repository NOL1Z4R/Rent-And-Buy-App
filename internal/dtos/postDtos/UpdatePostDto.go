package postDtos

import "Rent-And-Buy-App/internal/dtos/productDtos"

type UpdatePostDto struct {
	SellerId uint                         `json:"SellerId" binding:"required"`
	Product  productDtos.UpdateProductDto `json:"Product" binding:"required"`
	Price    float32                      `json:"Price" binding:"required"`
}
