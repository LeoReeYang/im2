package store

import "github.com/LeoReeYang/im2/models"

type MessageStore interface {
	Get() *models.Message
	Put(msg *models.Message) error
}
