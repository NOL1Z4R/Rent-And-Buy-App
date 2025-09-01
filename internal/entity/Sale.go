package entity

import (
	"gorm.io/gorm"
	"time"
)

type Sale struct {
	gorm.Model
	CustomerId uint      `json:"CustomerId" gorm:"not null" binding:"required"`
	PostId     uint      `json:"PostId" gorm:"not null" binding:"required"`
	SaleDate   time.Time `json:"SaleDate" gorm:"not null" binding:"required"`
}
