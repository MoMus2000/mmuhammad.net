package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Not using automigrate, since we have the sql script for the table
type Monitor struct {
	gorm.Model
	Metric string `gorm:"primaryKey"`
	Value  string
	Date   string
}

type MonitorService struct {
	db *gorm.DB
}

func (ms *MonitorService) AutoMigrate() error {
	err := ms.db.AutoMigrate(&Monitor{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewMonitorService(db *gorm.DB) *MonitorService {
	return &MonitorService{
		db: db,
	}
}

func (ms *MonitorService) UsdToPkr() ([][]string, error) {
	currentTime := time.Now().AddDate(0, 0, -5).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ?", "OPEN_USD", "CLOSE_USD").
		Where("date >= ?", currentTime).
		Find(&monitor).Error
	if err != nil {
		return nil, err
	}
	for _, result := range monitor {
		monitorString = append(monitorString, []string{
			result.Metric,
			result.Value,
			result.Date,
		})
	}

	return monitorString, nil
}

func (ms *MonitorService) MetalPrices() ([][]string, error) {
	currentTime := time.Now().AddDate(0, 0, -5).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ?", "CHINA_HOT_ROLL", "TURKEY_REBAR", "TURKEY_SC").
		Where("date >= ?", currentTime).
		Find(&monitor).
		Error

	if err != nil {
		return nil, err
	}
	for _, result := range monitor {
		monitorString = append(monitorString, []string{
			result.Metric,
			result.Value,
			result.Date,
		})
	}

	return monitorString, nil
}

func (ms *MonitorService) OilPrices() ([][]string, error) {
	currentTime := time.Now().AddDate(0, 0, -5).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ?", "WTIOIL", "BRENTOIL").
		Where("date >= ?", currentTime).
		Find(&monitor).
		Error

	if err != nil {
		return nil, err
	}
	for _, result := range monitor {
		monitorString = append(monitorString, []string{
			result.Metric,
			result.Value,
			result.Date,
		})
	}

	return monitorString, nil
}

func (ms *MonitorService) BasementRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "BASEMENT_MEAN", "BASEMENT_MIN", "BASEMENT_MEDIAN", "BASEMENT_MAX", "BASEMENT_LIKELY_PRICE").
		Where("date >= ?", currentTime).
		Find(&monitor).
		Error
	if err != nil {
		return nil, err
	}
	for _, result := range monitor {
		monitorString = append(monitorString, []string{
			result.Metric,
			result.Value,
			result.Date,
		})
	}

	return monitorString, nil
}

func (ms *MonitorService) ApartmentRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "APARTMENT_MEAN", "APARTMENT_MIN", "APARTMENT_MEDIAN", "APARTMENT_MAX", "APARTMENT_LIKELY_PRICE").
		Where("date >= ?", currentTime).
		Find(&monitor).
		Error
	if err != nil {
		return nil, err
	}
	for _, result := range monitor {
		monitorString = append(monitorString, []string{
			result.Metric,
			result.Value,
			result.Date,
		})
	}

	return monitorString, nil
}
