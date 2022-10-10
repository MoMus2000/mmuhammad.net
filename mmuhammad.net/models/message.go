package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Username string
	Message  string
}

type MessageService struct {
	db *gorm.DB
}

func (ms *MessageService) AutoMigrate() error {
	err := ms.db.AutoMigrate(&Message{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewMessageService(db *gorm.DB) *MessageService {
	return &MessageService{
		db: db,
	}
}

func (ms *MessageService) GetRecentMessages() ([][]string, error) {
	messages := []Message{}
	messageString := [][]string{}
	err := ms.db.Order("created_at DESC").Limit(5).Find(&messages).Error

	for _, result := range messages {
		messageString = append(messageString, []string{
			result.Username,
			result.Message,
		})
	}

	return messageString, err
}

func (ms *MessageService) WriteMessage(message *Message) error {
	return ms.db.Create(&message).Error
}
