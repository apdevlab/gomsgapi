package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Message data model
type Message struct {
	ID        uuid.UUID
	Message   string
	CreatedAt time.Time
}

// NewMessage initialize new instance of Message
func NewMessage(message string) *Message {
	return &Message{
		ID:        uuid.NewV1(),
		Message:   message,
		CreatedAt: time.Now().UTC(),
	}
}
