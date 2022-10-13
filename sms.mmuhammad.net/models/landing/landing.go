package landing

import (
	"time"

	"github.com/jinzhu/gorm"
)

type LandingContactForm struct {
	Name    string    `gorm:"not_null"`
	Email   string    `gorm:"not_null"`
	Message string    `gorm:"not_null"`
	Subject string    `gorm:"not_null"`
	Time    time.Time `gorm:"not_null"`
}

type LandingService struct {
	db *gorm.DB
}

func NewLandingService(db *gorm.DB) *LandingService {
	return &LandingService{
		db: db,
	}
}

func (ls *LandingService) AutoMigrate() error {
	return ls.db.AutoMigrate(&LandingContactForm{}).Error
}

func (ls *LandingService) SaveContactInfo(ld interface{}) error {
	return ls.db.Begin().Create(ld).Commit().Error
}
