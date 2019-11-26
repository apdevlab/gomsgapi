package services

import (
	"errors"
	"sync"

	"gomsgapi/modules/messages/dto"
	"gomsgapi/modules/messages/models"
	"gomsgapi/modules/messages/repositories"

	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
)

var (
	messageSvcLock sync.Once
	messageSvc     MessageService
)

// MessageService interface
type MessageService interface {
	Create(message dto.CreateMessageRequest) (*dto.Message, error)
	GetByID(id uuid.UUID) (*dto.Message, error)
	GetAll() ([]dto.Message, error)
}

type messageService struct {
	repository repositories.MessageRepository
}

// NewMessageService initialize new message service
func NewMessageService(repository repositories.MessageRepository) MessageService {
	messageSvcLock.Do(func() {
		messageSvc = &messageService{
			repository: repository,
		}
	})

	return messageSvc
}

func (svc messageService) Create(message dto.CreateMessageRequest) (*dto.Message, error) {
	model := models.NewMessage(message.Message)
	if err := svc.repository.Save(*model); err != nil {
		return nil, err
	}

	var result dto.Message
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc messageService) GetByID(id uuid.UUID) (*dto.Message, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid id")
	}

	model, err := svc.repository.Get(id)
	if err != nil {
		return nil, err
	}

	var result dto.Message
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc messageService) GetAll() ([]dto.Message, error) {
	models, err := svc.repository.GetAll()
	if err != nil {
		return nil, err
	}

	var results []dto.Message
	if err := copier.Copy(&results, &models); err != nil {
		return nil, err
	}

	return results, nil
}
