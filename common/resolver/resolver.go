package resolver

import (
	"gomsgapi/infra/websocket"
	inmemoryMessageRepositories "gomsgapi/modules/messages/repositories/inmemory"
	messageServices "gomsgapi/modules/messages/services"
)

// Resolver struct
type Resolver struct {
	MessageService messageServices.MessageService
	Websocket      websocket.Websocket
}

// NewResolver initialize new resolver object
func NewResolver() *Resolver {
	// storages
	messageStorage := inmemoryMessageRepositories.NewMessageStorage()

	// repositories
	messageRepository := inmemoryMessageRepositories.NewMessageRepositoryInMemory(messageStorage)

	// services
	messageService := messageServices.NewMessageService(messageRepository)

	// infras
	websocket := websocket.NewWebsocket()

	return &Resolver{
		MessageService: messageService,
		Websocket:      websocket,
	}
}
