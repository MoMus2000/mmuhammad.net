package models

import (
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
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Limit(10).Order("created_at ASC").Where("metric = ? OR metric = ?", "OPEN_USD", "CLOSE_USD").Find(&monitor).Error
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
