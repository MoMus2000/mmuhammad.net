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
	currentTime := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
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
	currentTime := time.Now().AddDate(0, 0, -15).Format("2006-01-02")
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
	currentTime := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
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

func (ms *MonitorService) SPYRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ?", "CLOSE_SPY").
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

func (ms *MonitorService) SPYRegimeProbs() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ?", "LOW_VOL_PROB_SPY", "MED_VOL_PROB_SPY", "HIGH_VOL_PROB_SPY").
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

func (ms *MonitorService) CADHousingRegime() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ?", "LOW_VOL_PROB_XRE.TO", "MED_VOL_PROB_XRE.TO", "HIGH_VOL_PROB_XRE.TO").
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

func (ms *MonitorService) DurhanApartmentRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_DURHAM_APARTMENT_MEAN", "TWO_BD_DURHAM_APARTMENT_MIN", "TWO_BD_DURHAM_APARTMENT_MEDIAN", "TWO_BD_DURHAM_APARTMENT_MAX", "TWO_BD_DURHAM_APARTMENT_LIKELY_PRICE").
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

func (ms *MonitorService) StCatharinesApartmentRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_ST_CATHARINES_APARTMENT_MEAN", "TWO_BD_ST_CATHARINES_APARTMENT_MIN", "TWO_BD_ST_CATHARINES_APARTMENT_MEDIAN", "TWO_BD_ST_CATHARINES_APARTMENT_MAX", "TWO_BD_ST_CATHARINES_APARTMENT_LIKELY_PRICE").
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

func (ms *MonitorService) HamiltonApartmentRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_HAMILTON_APARTMENT_MEAN", "TWO_BD_HAMILTON_APARTMENT_MIN", "TWO_BD_HAMILTON_APARTMENT_MEDIAN", "TWO_BD_HAMILTON_APARTMENT_MAX", "TWO_BD_HAMILTON_APARTMENT_LIKELY_PRICE").
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

func (ms *MonitorService) WindsorApartmentRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_WINDSOR_APARTMENT_MEAN", "TWO_BD_WINDSOR_APARTMENT_MIN", "TWO_BD_WINDSOR_APARTMENT_MEDIAN", "TWO_BD_WINDSOR_APARTMENT_MAX", "TWO_BD_WINDSOR_LIKELY_PRICE").
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

func (ms *MonitorService) WindsorBasementRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_WINDSOR_BASEMENT_MEAN", "TWO_BD_WINDSOR_BASEMENT_MIN", "TWO_BD_WINDSOR_BASEMENT_MEDIAN", "TWO_BD_WINDSOR_BASEMENT_MAX", "TWO_BD_WINDSOR_BASEMENT_LIKELY_PRICE").
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

func (ms *MonitorService) HamiltonBasementRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_HAMILTON_BASEMENT_MEAN", "TWO_BD_HAMILTON_BASEMENT_MIN", "TWO_BD_HAMILTON_BASEMENT_MEDIAN", "TWO_BD_HAMILTON_BASEMENT_MAX", "TWO_BD_HAMILTON_BASEMENT_LIKELY_PRICE").
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

func (ms *MonitorService) DurhamBasementRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_DURHAM_BASEMENT_MEAN", "TWO_BD_DURHAM_BASEMENT_MIN", "TWO_BD_DURHAM_BASEMENT_MEDIAN", "TWO_BD_DURHAM_BASEMENT_MAX", "TWO_BD_DURHAM_BASEMENT_LIKELY_PRICE").
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

func (ms *MonitorService) StCatharinesBasementRates() ([][]string, error) {
	currentTime := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	monitor := []Monitor{}
	monitorString := [][]string{}
	err := ms.db.Order("created_at ASC").
		Where("metric = ? OR metric = ? OR metric = ? OR metric = ? OR metric = ?", "TWO_BD_ST_CATHARINES_BASEMENT_MEAN", "TWO_BD_ST_CATHARINES_BASEMENT_MIN", "TWO_BD_ST_CATHARINES_BASEMENT_MEDIAN", "TWO_BD_ST_CATHARINES_BASEMENT_MAX", "TWO_BD_ST_CATHARINES_BASEMENT_LIKELY_PRICE").
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
