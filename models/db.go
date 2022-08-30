package models

import "github.com/jinzhu/gorm"

func NewDataBaseConnection(connectionInfo string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, err
}
