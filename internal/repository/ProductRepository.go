package repository

import (
	"Rent-And-Buy-App/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) GetAll() ([]*entity.Product, error) {
	var products []*entity.Product
	err := pr.db.Find(&products)
	return products, err.Error
}

func (pr *ProductRepository) GetById(id uint) (*entity.Product, error) {
	var product *entity.Product
	err := pr.db.Where("id = ?", id).Find(&product).Error
	return product, err
}

func (pr *ProductRepository) Create(product *entity.Product) error {
	err := pr.db.Create(&product).Error
	return err
}

func (pr *ProductRepository) Update(product *entity.Product) error {
	err := pr.db.Save(&product).Error
	return err
}

func (pr *ProductRepository) Delete(product *entity.Product) error {
	err := pr.db.Delete(&product).Error
	return err
}
