package models

import (
	"errors"
	"fmt"

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

func NewAdminService(db *gorm.DB) *AdminService {
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
		return err
	}
	admin.Password = string(hashedBytes)
	return a.db.Create(admin).Error
}

func (a *AdminService) UpdateChangesFromEdit(post *Post, id string) error {
	fmt.Println(post)
	if id == "" {
		return errors.New("Id has to be provided")
	}
	if post.Imgur_URL != "" {
		err := a.db.Model(&post).Where("id = ?", id).
			Update("imgur_url", post.Imgur_URL).Error
		if err != nil {
			return err
		}
	}
	if post.Topic != "" {
		err := a.db.Model(&post).Where("id = ?", id).
			Update("topic", post.Topic).Error
		if err != nil {
			return err
		}
	}
	if post.Summary != "" {
		err := a.db.Model(&post).Where("id = ?", id).
			Update("summary", post.Summary).Error
		if err != nil {
			return err
		}
	}
	if post.Content != "" {
		err := a.db.Model(&post).Where("id = ?", id).
			Update("content", post.Content).Error
		if err != nil {
			return err
		}
	}
	return nil
}
