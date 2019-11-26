package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Message dto
type Message struct {
	ID        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateMessageRequest dto
type CreateMessageRequest struct {
	Message string `json:"message"`
}
