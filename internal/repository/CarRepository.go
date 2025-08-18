package repository

import (
	"Rent-And-Buy-App/internal/entity"
	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (cr *CarRepository) Create(car *entity.Car) error {

	return cr.db.Create(&car).Error
}

func (cr *CarRepository) GetAll() ([]*entity.Car, error) {
	var cars []*entity.Car
	result := cr.db.Find(&cars)
	return cars, result.Error
}

func (cr *CarRepository) GetById(id uint) (*entity.Car, error) {
	var car entity.Car
	result := cr.db.Where("id = ?", id).Find(&car)
	return &car, result.Error
}

func (cr *CarRepository) GetByPlate(plate string) (*entity.Car, error) {
	var car entity.Car
	result := cr.db.Where("plate = ?", plate).Find(&car)
	return &car, result.Error
}

func (cr *CarRepository) Update(car *entity.Car) error {
	return cr.db.Save(&car).Error
}

func (cr *CarRepository) Delete(car *entity.Car) error {
	return cr.db.Delete(&car).Error
}
