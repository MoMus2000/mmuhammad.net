package model_sms

import (
	"github.com/jinzhu/gorm"
)

type SmsMetrics struct {
	gorm.Model
	Email       string `gorm:"not_null"`
	MetricName  string `gorm:"not_null"`
	MetricValue string `gorm:"not_null"`
}

type SmsMetricService struct {
	db *gorm.DB
}

func NewSmsMetricService(db *gorm.DB) *SmsMetricService {
	return &SmsMetricService{db}
}

func (sms *SmsMetricService) AutoMigrate() error {
	return sms.db.AutoMigrate(&SmsMetrics{}).Error
}

func (sms *SmsMetricService) GetTotalMessages(userId string) (float32, error) {
	count := 0
	err := sms.db.Model(&SmsMetrics{}).Where("email = ?", userId).Count(&count).Error
	return float32(count), err
}

func (sms *SmsMetricService) GetTotalPrices(userId string) (float32, error) {
	return 10.0, nil
}

func (sms *SmsMetricService) GetTodayPrices(userId string) (float32, error) {
	return 10.0, nil
}

func (sms *SmsMetricService) GetTodayMessages(userId string) (float32, error) {
	count := 0
	err := sms.db.Model(&SmsMetrics{}).Where("email = ?", userId).Count(&count).Error
	return float32(count), err
}
