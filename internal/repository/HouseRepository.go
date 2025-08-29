package repository

import (
	"Rent-And-Buy-App/internal/entity"
	"gorm.io/gorm"
)

type HouseRepository struct {
	db *gorm.DB
}

func NewHouseRepository(db *gorm.DB) *HouseRepository {
	return &HouseRepository{db: db}
}

func (hr *HouseRepository) GetAll() ([]*entity.House, error) {
	var houses []*entity.House
	result := hr.db.Find(&houses)
	return houses, result.Error
}

func (hr *HouseRepository) GetById(id uint) (*entity.House, error) {
	var house *entity.House
	result := hr.db.Where("id = ?", id).Find(&house)
	return house, result.Error
}

func (hr *HouseRepository) Create(house *entity.House) error {
	return hr.db.Create(&house).Error
}
func (hr *HouseRepository) Update(house *entity.House) error {
	result := hr.db.Save(house)
	return result.Error
}

func (hr *HouseRepository) Delete(house *entity.House) error {
	result := hr.db.Delete(&house)
	return result.Error
}
