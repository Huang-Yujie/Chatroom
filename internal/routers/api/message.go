package api

import "github.com/gin-gonic/gin"

type Message struct{}

func NewMessage() Message {
	return Message{}
}

func (m Message) Send(c *gin.Context) {}

func (m Message) Broadcast(c *gin.Context) {}
