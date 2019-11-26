package inmemory

import (
	"gomsgapi/modules/messages/models"
)

// MessageStorage to store message in memory
type MessageStorage struct {
	Messages []models.Message
}

// NewMessageStorage initialize new message storage object
func NewMessageStorage() *MessageStorage {
	return &MessageStorage{
		Messages: []models.Message{},
	}
}
