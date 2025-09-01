package repository

import (
	"Rent-And-Buy-App/internal/entity"
	"gorm.io/gorm"
)

type SaleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

func (sr *SaleRepository) GetAll() ([]*entity.Sale, error) {
	var sales []*entity.Sale
	result := sr.db.Find(&sales)
	if result.Error != nil {
		return nil, result.Error
	}
	return sales, nil
}

func (sr *SaleRepository) GetById(id uint) (*entity.Sale, error) {
	var sale *entity.Sale
	result := sr.db.First(&sale, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return sale, nil
}

func (sr *SaleRepository) Create(sale *entity.Sale) (*entity.Sale, error) {
	result := sr.db.Create(&sale)
	return sale, result.Error
}

func (sr *SaleRepository) Update(sale *entity.Sale) (*entity.Sale, error) {
	result := sr.db.Save(&sale)
	return sale, result.Error
}

func (sr *SaleRepository) Delete(sale *entity.Sale) error {
	result := sr.db.Delete(&sale)
	return result.Error
}
