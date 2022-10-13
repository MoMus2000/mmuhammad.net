package db

import ("github.com/jinzhu/gorm"
_ "github.com/mattn/go-sqlite3"
)


func NewDbConnection(path string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}
