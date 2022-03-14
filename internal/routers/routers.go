package routers

import (
	"net/http"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/Huang-Yujie/Chatroom/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSettings.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	}
	user := api.NewUser()
	message := api.NewMessage()
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "test"})
		})

		apiGroup.POST("/register", user.Register)
		apiGroup.POST("/login", user.Login)
		apiGroup.POST("/send", message.Send)
	}
	return r
}
