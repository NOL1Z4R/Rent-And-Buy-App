package entity

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Brand    string `json:"Brand" binding:"required"`
	CarModel string `json:"CarModel" binding:"required"`
	Year     uint   `json:"Year" binding:"required"`
	Milage   uint   `json:"Milage" binding:"required"`
}
