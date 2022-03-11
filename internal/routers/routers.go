package routers

import (
	"net/http"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSettings.Runmode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	}
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "test"})
		})
	}
	return r
}
