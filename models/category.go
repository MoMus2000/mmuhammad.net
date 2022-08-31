package models

import "github.com/jinzhu/gorm"

type Category struct {
	*gorm.Model
	CategoryName    string
	CategorySummary string
	CreationDate    string
	Imgur_URL       string
}

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (cs *CategoryService) AutoMigrate() error {
	return cs.db.AutoMigrate(&Category{}).Error
}

func (cs *CategoryService) Create(cat *Category) error {
	return cs.db.Create(cat).Error
}
