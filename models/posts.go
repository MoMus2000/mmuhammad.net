package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(connectionInfo string) *PostService {
	db, err := gorm.Open("sqlite3", connectionInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return &PostService{
		db: db,
	}
}

func (p *PostService) AutoMigrate() error {
	err := p.db.AutoMigrate(&Post{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostService) Create(post *Post) error {
	return p.db.Create(post).Error
}

type Post struct {
	gorm.Model
	Topic   string `gorm:"not null"`
	Content string `gorm:"not null"`
	Date    string `gorm:"not null"`
}
