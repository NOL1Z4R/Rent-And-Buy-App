package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CarId   *uint `json:"CarId"`
	HouseId *uint `json:"HouseId"`
}
