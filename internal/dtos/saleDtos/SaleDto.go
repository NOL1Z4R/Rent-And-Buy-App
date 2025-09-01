package saleDtos

import (
	"Rent-And-Buy-App/internal/dtos/postDtos"
	"time"
)

type SaleDto struct {
	ID         uint             `json:"id" gorm:"not null" binding:"required"`
	CustomerId uint             `json:"CustomerId" binding:"required"`
	Post       postDtos.PostDto `json:"Post" binding:"required"`
	SaleDate   time.Time        `json:"SaleDate" binding:"required"`
}
