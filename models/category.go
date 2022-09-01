package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

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

func (cs *CategoryService) GetAllCategories() ([][]string, error) {
	categories := []Category{}
	categoryString := [][]string{}
	results := cs.db.Find(&categories).Order("Date DESC")
	for _, cat := range categories {
		categoryString = append(categoryString, []string{
			cat.CategoryName,
			cat.CategorySummary,
			cat.Imgur_URL,
			cat.CreationDate,
			strconv.FormatUint(uint64(cat.ID), 10),
		})
	}
	err := results.Error

	if err != nil {
		return nil, err
	}

	return categoryString, nil
}
