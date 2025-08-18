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

func (hb *HouseRepository) GetAll() ([]*entity.House, error) {
	var houses []*entity.House
	result := hb.db.Find(&houses)
	return houses, result.Error
}

func (hb *HouseRepository) GetById(id uint) (*entity.House, error) {
	var house *entity.House
	result := hb.db.Where("id = ?", id).Find(&house)
	return house, result.Error
}

func (hb *HouseRepository) Create(house *entity.House) error {
	result := hb.db.Create(house)
	return result.Error
}
func (hb *HouseRepository) Update(house *entity.House) error {
	result := hb.db.Save(house)
	return result.Error
}

func (hb *HouseRepository) Delete(house *entity.House) error {
	result := hb.db.Delete(&house)
	return result.Error
}
