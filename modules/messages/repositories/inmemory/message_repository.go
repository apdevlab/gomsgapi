package inmemory

import (
	"errors"
	"sync"

	"gomsgapi/modules/messages/models"
	"gomsgapi/modules/messages/repositories"

	uuid "github.com/satori/go.uuid"
)

var (
	messageRepoLock sync.Once
	messageRepo     repositories.MessageRepository
)

type messageRepository struct {
	storage *MessageStorage
}

// NewMessageRepositoryInMemory initialize new message repository object
func NewMessageRepositoryInMemory(storage *MessageStorage) repositories.MessageRepository {
	messageRepoLock.Do(func() {
		messageRepo = &messageRepository{
			storage: storage,
		}
	})

	return messageRepo
}

func (repo *messageRepository) Get(id uuid.UUID) (*models.Message, error) {
	message := models.Message{}
	for _, item := range repo.storage.Messages {
		if item.ID == id {
			message = item
			break
		}
	}

	if message == (models.Message{}) {
		return nil, errors.New("not found")
	}

	return &message, nil
}

func (repo *messageRepository) GetAll() ([]models.Message, error) {
	messages := []models.Message{}
	for _, item := range repo.storage.Messages {
		messages = append(messages, item)
	}

	return messages, nil
}

func (repo *messageRepository) Save(message models.Message) error {
	found := false
	for idx, msg := range repo.storage.Messages {
		if msg.ID == message.ID {
			found = true
			repo.storage.Messages[idx] = message
			break
		}
	}

	if !found {
		repo.storage.Messages = append(repo.storage.Messages, message)
	}

	return nil
}
