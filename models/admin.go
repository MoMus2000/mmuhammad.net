package models

import "github.com/jinzhu/gorm"

type AdminService struct {
	db *gorm.DB
}

type Admin struct {
	Email    string `gorm:"not_null;unique_index"`
	Password string `gorm:"not_null"`
}

// I need token functions
// I need hashing functions
// I need password lookup
// I need email lookup
