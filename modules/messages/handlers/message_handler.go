package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gomsgapi/common/resolver"
	"gomsgapi/modules/messages/dto"

	"github.com/gin-gonic/gin"
)

// MessageHandler struct
type MessageHandler struct {
	resolver *resolver.Resolver
}

// NewMessageHandler initialize new message handler object
func NewMessageHandler(resolver *resolver.Resolver) *MessageHandler {
	handler := &MessageHandler{
		resolver: resolver,
	}

	resolver.Websocket.RegisterListener(handler.WsListener)
	return handler
}

// Submit new message
func (h *MessageHandler) Submit(ctx *gin.Context) {
	var request dto.CreateMessageRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if request.Message == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "message is required"})
		return
	}

	message, err := h.resolver.MessageService.Create(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.resolver.Websocket.Broadcast(fmt.Sprintf("[%s] %s", message.CreatedAt.Format(time.RFC3339), message.Message)); err != nil {
		log.Printf("failed to broadcast message: %s", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": message,
	})
}

// GetAll godoc
func (h *MessageHandler) GetAll(ctx *gin.Context) {
	messages, err := h.resolver.MessageService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}

// HandleWs godoc
func (h *MessageHandler) HandleWs(ctx *gin.Context) {
	err := h.resolver.Websocket.Open(ctx.Writer, ctx.Request)
	if err != nil {
		log.Printf("websocket terminated with error: %s", err)
	}
}

// WsListener godoc
func (h *MessageHandler) WsListener(message string) {
	log.Printf("receiving message via ws: %s", message)
}
