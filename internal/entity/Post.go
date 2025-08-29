package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	SellerId  uint    `json:"Seller" binding:"required"`
	ProductId uint    `json:"ProductId" binding:"required"`
	Price     float32 `json:"Price" binding:"required"`
}
