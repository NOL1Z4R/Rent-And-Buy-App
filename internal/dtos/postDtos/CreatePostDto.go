package postDtos

import "Rent-And-Buy-App/internal/dtos/productDtos"

type CreatePostDto struct {
	SellerId uint                         `json:"SellerId" binding:"required"`
	Product  productDtos.CreateProductDto `json:"Product" binding:"required"`
	Price    float32                      `json:"Price" binding:"required"`
}
