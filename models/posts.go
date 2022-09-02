package models

import (
	"fmt"
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
			post.CategoryId,
		})
	}
	err := results.Error

	if err != nil {
		return nil, err
	}

	return postString, nil
}

func (ps *PostService) GetAllPostByCategory(cid string) ([][]string, error) {
	fmt.Println("We are here")
	posts := []Post{}
	postString := [][]string{}
	results := ps.db.Where("category_id = ?", cid).Find(&posts)
	for _, post := range posts {
		postString = append(postString, []string{
			post.Topic,
			post.Summary,
			post.Imgur_URL,
			post.Date,
			strconv.FormatUint(uint64(post.ID), 10),
			post.CategoryId,
		})
	}
	err := results.Error

	if err != nil {
		return nil, err
	}

	fmt.Println(postString)
	return postString, nil
}

type Post struct {
	gorm.Model
	Topic      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	Imgur_URL  string `gorm:"not null"`
	Summary    string `gorm:"not null"`
	Date       string `gorm:""`
	CategoryId string `gorm:"not null"`
}
