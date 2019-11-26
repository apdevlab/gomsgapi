package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to set websocket upgrade: %s", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("failed to read message from websocket: %s", err)
			break
		}

		conn.WriteMessage(t, msg)
	}
}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "PONG!")
	})

	api := router.Group("/api")
	{
		api.POST("/message", func(ctx *gin.Context) {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Not implemented yet"})
		})
		api.GET("/message", func(ctx *gin.Context) {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Not implemented yet"})
		})
		api.GET("/message/ws", func(ctx *gin.Context) {
			wshandler(ctx.Writer, ctx.Request)
		})
	}

	router.Run(":8080")
}
