package inmemory

import (
	"testing"
	"time"

	"gomsgapi/modules/messages/models"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	id := uuid.NewV1()
	created := time.Now().UTC()
	message := "Test message"

	storage := NewMessageStorage()
	storage.Messages = []models.Message{
		models.Message{ID: id, Message: message, CreatedAt: created},
	}

	tests := []struct {
		name    string
		id      uuid.UUID
		wantErr bool
		result  interface{}
	}{
		{
			name:    "message not found",
			id:      uuid.NewV1(),
			wantErr: true,
		},
		{
			name:   "message found",
			id:     id,
			result: &models.Message{ID: id, Message: message, CreatedAt: created},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := NewMessageRepositoryInMemory(storage)
			result, err := repo.Get(test.id)

			if test.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, "not found")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, result)
			}
		})
	}
}

func TestGetByAll(t *testing.T) {
	message := "Test message"

	storage := NewMessageStorage()
	storage.Messages = []models.Message{
		models.Message{Message: message},
	}

	tests := []struct {
		name   string
		result []models.Message
	}{
		{
			name: "messages found",
			result: []models.Message{
				models.Message{Message: message},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := NewMessageRepositoryInMemory(storage)
			result, err := repo.GetAll()

			assert.NoError(t, err)
			assert.Equal(t, len(test.result), len(result))
			assert.Equal(t, test.result[0].Message, result[0].Message)
		})
	}
}

func TestSave(t *testing.T) {
	id := uuid.NewV1()
	storage := NewMessageStorage()

	tests := []struct {
		name    string
		message models.Message
		result  models.Message
	}{
		{
			name:    "test insert",
			message: models.Message{ID: id, Message: "test message 1"},
			result:  models.Message{ID: id, Message: "test message 1"},
		},
		{
			name:    "test update",
			message: models.Message{ID: id, Message: "test message update"},
			result:  models.Message{ID: id, Message: "test message update"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := NewMessageRepositoryInMemory(storage)
			repo.Save(test.message)
			result, _ := repo.Get(test.message.ID)

			assert.Equal(t, test.result.ID, result.ID)
			assert.Equal(t, test.result.Message, result.Message)
		})
	}
}
