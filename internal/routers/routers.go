package routers

import (
	"net/http"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/Huang-Yujie/Chatroom/internal/middleware"
	"github.com/Huang-Yujie/Chatroom/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSettings.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	}
	user := api.NewUser()
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/register", user.Register)
		apiGroup.POST("/login", user.Login)
	}
	wsGroup := r.Group("/ws")
	wsGroup.Use(middleware.JWT())
	{
		wsGroup.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "test"})
		})
		wsGroup.GET("/", WebsocketHandler)
	}
	return r
}
