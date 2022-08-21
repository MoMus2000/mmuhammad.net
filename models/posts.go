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

func (p *PostService) GetAllPost() ([]string, error) {
	posts := []Post{}
	postString := []string{}
	results := p.db.Find(&posts)
	for _, post := range posts {
		postString = append(postString, string(post.Topic))
		postString = append(postString, string(post.Content))
		postString = append(postString, string(post.Date))
		postString = append(postString, string(post.Summary))
	}
	err := results.Error

	if err != nil {
		return nil, err
	}

	return postString, nil
}

type Post struct {
	gorm.Model
	Topic   string `gorm:"not null"`
	Content string `gorm:"not null"`
	Date    string `gorm:"not null"`
	Summary string `gorm:"not null"`
}
