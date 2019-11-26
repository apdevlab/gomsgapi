package repositories

import (
	"gomsgapi/modules/messages/models"

	uuid "github.com/satori/go.uuid"
)

// MessageRepository interface
type MessageRepository interface {
	Get(id uuid.UUID) (*models.Message, error)
	GetAll() ([]models.Message, error)
	Save(message models.Message) error
}
