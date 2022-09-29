package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const Fmbpepper = "pepper"

type FmbService struct {
	db *gorm.DB
}

type Fmb struct {
	Email    string `gorm:"not_null;unique_index"`
	Password string `gorm:"not_null"`
}

func (fm *FmbService) AutoMigrate() error {
	return fm.db.AutoMigrate(&Fmb{}).Error
}

func NewFmbService(db *gorm.DB) *FmbService {
	return &FmbService{db: db}
}

func (a *FmbService) ByEmail(fm *Fmb) (*Fmb, error) {
	fmb := Fmb{}
	err := a.db.Where("email = ?", fm.Email).First(&fmb).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(fmb.Password), []byte(fm.Password+Fmbpepper))
	if err != nil {
		return nil, err
	}
	return &fmb, nil
}

func (a *FmbService) Create(fmb *Fmb) error {
	passwordBytes := []byte(fmb.Password + Fmbpepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmb.Password = string(hashedBytes)
	return a.db.Create(fmb).Error
}
