package models

import (
	"errors"
	"fmt"
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
	results := cs.db.Order("created_at ASC").Find(&categories)
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

func (cs *CategoryService) DeleteCategory(id uint64) error {
	return cs.db.Delete(&Category{}, id).Error
}

func (cs *CategoryService) UpdateChangesCategoryFromEdit(category *Category, id string) error {
	fmt.Println(category)
	if id == "" {
		return errors.New("Id has to be provided")
	}
	if category.Imgur_URL != "" {
		err := cs.db.Model(&category).Where("id = ?", id).
			Update("imgur_url", category.Imgur_URL).Error
		if err != nil {
			return err
		}
	}
	if category.CategoryName != "" {
		err := cs.db.Model(&category).Where("id = ?", id).
			Update("category_name", category.CategoryName).Error
		if err != nil {
			return err
		}
	}
	if category.CategorySummary != "" {
		err := cs.db.Model(&category).Where("id = ?", id).
			Update("category_summary", category.CategorySummary).Error
		if err != nil {
			return err
		}
	}
	return nil
}
