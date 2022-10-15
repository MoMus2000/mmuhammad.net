package model_auth

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	db *gorm.DB
}

func (auth *AuthService) ByEmail(Email string, Password string) (*User, error) {
	user := User{}
	err := auth.db.Where("email = ?", Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password+pepper))
	if err != nil {
		return nil, err
	}
	return &user, nil
}
