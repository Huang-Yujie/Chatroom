package api

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Register(c *gin.Context) {}

func (u User) Login(c *gin.Context) {}
