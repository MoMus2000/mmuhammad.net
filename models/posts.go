package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		db: db,
	}
}

func (ps *PostService) AutoMigrate() error {
	err := ps.db.AutoMigrate(&Post{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) CreatePost(p *Post) error {
	return ps.db.Create(p).Error
}

func (ps *PostService) DeletePost(id uint) error {
	return ps.db.Delete(&Post{}, id).Error
}

func (ps *PostService) GetPost(id uint) *Post {
	post := Post{}
	ps.db.First(&post, id)
	return &post
}

func (ps *PostService) GetAllPost() ([][]string, error) {
	posts := []Post{}
	postString := [][]string{}
	results := ps.db.Find(&posts).Order("Date DESC")
	for _, post := range posts {
		postString = append(postString, []string{
			post.Topic,
			post.Summary,
			post.Imgur_URL,
			post.Date,
			strconv.FormatUint(uint64(post.ID), 10),
		})
	}
	err := results.Error

	if err != nil {
		return nil, err
	}

	return postString, nil
}

type Post struct {
	gorm.Model
	Topic     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	Imgur_URL string `gorm:"not null"`
	Summary   string `gorm:"not null"`
	Date      string `gorm:""`
}
