package models

import "github.com/jinzhu/gorm"

type AdminService struct {
	db *gorm.DB
}

type Admin struct {
	Email    string `gorm:"not_null;unique_index"`
	Password string `gorm:"not_null"`
}

func (a *AdminService) AutoMigrate() {
	a.db.AutoMigrate(&Admin{})
}

func (a *AdminService) ByEmail(email string) (*Admin, error) {
	admin := Admin{}
	err := a.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// I need token functions
// I need hashing functions
// I need password lookup
// I need email lookup
