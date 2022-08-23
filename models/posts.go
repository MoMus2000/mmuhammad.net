package models

import (
	"strconv"

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

func (ps *PostService) CreatePost(p *Post) error {
	return ps.db.Create(p).Error
}

func (ps *PostService) GetPost(id uint) *Post {
	post := Post{}
	ps.db.First(&post, id)
	return &post
}

func (p *PostService) GetAllPost() ([][]string, error) {
	posts := []Post{}
	postString := [][]string{}
	results := p.db.Find(&posts)
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
	Imgur_URL string `gorm:""`
	Summary   string `gorm:"not null"`
	Date      string `gorm:""`
}
