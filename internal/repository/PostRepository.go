package repository

import (
	"Rent-And-Buy-App/internal/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) GetAll() ([]*entity.Post, error) {
	var posts []*entity.Post
	err := pr.db.Find(&posts)
	if err.Error != nil {
		return nil, err.Error
	}
	return posts, nil
}

func (pr *PostRepository) GetById(id uint) (*entity.Post, error) {
	var post *entity.Post
	err := pr.db.Where("id = ?", id).Find(&post)
	if err.Error != nil {
		return nil, err.Error
	}
	return post, nil
}

func (pr *PostRepository) Create(post *entity.Post) error {
	return pr.db.Create(&post).Error
}

func (pr *PostRepository) Update(post *entity.Post) error {
	return pr.db.Save(&post).Error
}

func (pr *PostRepository) Delete(post *entity.Post) error {
	return pr.db.Delete(&post).Error
}

//TODO offer part can be in here
