package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const pepper = "pepper"

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

func NewAdminService(connectionInfo string) *AdminService {
	db, err := gorm.Open("sqlite3", connectionInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return &AdminService{db: db}
}

func (a *AdminService) ByEmail(ad *Admin) (*Admin, error) {
	admin := Admin{}
	err := a.db.Where("email = ?", ad.Email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(ad.Password+pepper))
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a *AdminService) Create(admin *Admin) error {
	passwordBytes := []byte(admin.Password + pepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	admin.Password = string(hashedBytes)
	return a.db.Create(admin).Error
}

// I need token functions
// I need hashing functions
// I need password lookup
// I need email lookup
