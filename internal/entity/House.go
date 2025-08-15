package entity

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Adress      string `json:"Adress" binding:"required"`
	SquareMeter string `json:"SquareMeter" binding:"required"`
	Year        uint   `json:"Year" binding:"required"`
}
