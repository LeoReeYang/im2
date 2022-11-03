package store

import "im/models"

type MessageStore interface {
	Get() *models.Message
	Put(msg *models.Message) error
}
