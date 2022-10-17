package model_sms

import (
	"github.com/jinzhu/gorm"
)

type SmsMetrics struct {
	Email       string `gorm:"unique_index; not_null"`
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
	return 80.0, nil
}

func (sms *SmsMetricService) GetTotalPrices(userId string) (float32, error) {
	return 10.0, nil
}

func (sms *SmsMetricService) GetTodayPrices(userId string) (float32, error) {
	return 10.0, nil
}

func (sms *SmsMetricService) GetTodayMessages(userId string) (float32, error) {
	return 31.0, nil
}
