package model_sms

import (
	"strconv"
	"time"

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
	results := []SmsMetrics{}
	query := sms.db.Where("email = ? AND (metric_name = ?)", userId, "SMS_SENT").Find(&results)
	totalMessages := 0.0
	for _, result := range results {
		f, err := strconv.ParseFloat(result.MetricValue, 32)
		if err != nil {
			continue
		}
		totalMessages += f
	}

	return float32(totalMessages), query.Error
}

func (sms *SmsMetricService) GetTotalPrices(userId string) (float32, error) {
	return 10.0, nil
}

func (sms *SmsMetricService) GetTodayPrices(userId string) (float32, error) {
	return 10.0, nil
}

func (sms *SmsMetricService) GetTodayMessages(userId string) (float32, error) {
	results := []SmsMetrics{}

	currentDay := time.Now()
	diff := 24 * time.Hour
	prevDay := currentDay.Add(-diff)

	query := sms.db.Where("email = ? AND (metric_name = ?) AND created_at >= ?", userId, "SMS_SENT", prevDay).Find(&results)
	totalMessages := 0.0

	for _, result := range results {
		f, err := strconv.ParseFloat(result.MetricValue, 32)
		if err != nil {
			continue
		}
		totalMessages += f
	}

	return float32(totalMessages), query.Error
}
