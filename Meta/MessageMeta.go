package Meta

import (
	"github.com/LeoReeYang/im2/models"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

type MessageMeta struct {
	db *gorm.DB
}

func NewMessageMeta() *MessageMeta {
	return &MessageMeta{
		db: models.GetDB(),
	}
}

func (s *MessageMeta) Get(name string) []models.Message {
	var messages []models.Message
	s.db.Where("recipient = ?", name).Find(&messages)
	return messages
}

func (s *MessageMeta) Put(msg *models.Message) error {
	err := s.db.Create(msg).Error
	if msg.ID == 0 {
		color.Red("Message store failed:", err)
	}
	return nil
}
