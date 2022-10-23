package model_auth

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const pepper = "pepper"

type AuthService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Email        string `gorm:"not_null;unique_index"`
	Password     string `gorm:"not_null;"`
	TwilioPhone  string `gorm:""`
	AccountId    string `gorm:""`
	AccountToken string `gorm:""`
}

func (ucs *AuthService) AutoMigrate() error {
	return ucs.db.AutoMigrate(&User{}).Error
}

func NewCreateUserService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (ucs *AuthService) CreateUserService(Email string, Password string) error {
	passwordBytes := []byte(Password + pepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := User{Email: Email, Password: string(hashedBytes)}
	return ucs.db.Create(&user).Error
}
