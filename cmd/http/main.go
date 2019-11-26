package main

import (
	"net/http"

	"gomsgapi/common/resolver"
	messageHandlers "gomsgapi/modules/messages/handlers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	messageHandler *messageHandlers.MessageHandler
)

func init() {
	resolver := resolver.NewResolver()

	messageHandler = messageHandlers.NewMessageHandler(resolver)
}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "PONG!")
	})

	api := router.Group("/api")
	{
		api.POST("/message", messageHandler.Submit)
		api.GET("/message", messageHandler.GetAll)
		api.GET("/message/ws", messageHandler.HandleWs)
	}

	router.Run(":8080")
}
