package store

import (
	"github.com/LeoReeYang/im2/models"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

type MessageStore interface {
	Get() *models.Message
	Put(msg *models.Message) error
}

type SimpleStore struct {
	db *gorm.DB
}

func NewSimpleStore() *SimpleStore {
	return &SimpleStore{
		db: models.GetDB(),
	}
}

func (s *SimpleStore) Get() *models.Message {
	return nil
}

func (s *SimpleStore) Put(msg *models.Message) error {
	err := s.db.Create(msg).Error
	if msg.ID == 0 {
		color.Red("Message store failed:", err)
	}
	return nil
}
