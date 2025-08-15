package repository

import (
	"Rent-And-Buy-App/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type userRepository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) error
	GetByEmail(email string) (*entity.User, error)
	GetById(id uint) (*entity.User, error)
	Update(user entity.User) error
	Delete(user entity.User) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User

	result := ur.db.Find(&users)

	return users, result.Error
}

func (ur *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := ur.db.Where("email=?", email).Find(&user)
	return &user, result.Error
}
func (ur *UserRepository) GetById(id uint) (*entity.User, error) {
	var user entity.User
	result := ur.db.Where("id=?", id).Find(&user)
	return &user, result.Error
}

func (ur *UserRepository) Create(user *entity.User) error {
	return ur.db.Create(&user).Error
}

func (ur *UserRepository) Update(user *entity.User) error {
	return ur.db.Save(&user).Error
}

func (ur *UserRepository) Delete(user *entity.User) error {
	return ur.db.Delete(&user).Error
}
