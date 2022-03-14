package api

import (
	"github.com/Huang-Yujie/Chatroom/internal/request"
	"github.com/Huang-Yujie/Chatroom/internal/service"
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
	"github.com/Huang-Yujie/Chatroom/pkg/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Register(c *gin.Context) {
	param := request.UserRegisterRequest{}
	response := response.NewResponse(c)
	if c.ShouldBindJSON(&param) != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.UserRegister(&param)
	if err != nil {
		response.ToErrorResponse(err)
	}
	response.ToResponse(user)
}

func (u User) Login(c *gin.Context) {
	param := request.UserLoginRequest{}
	response := response.NewResponse(c)
	if c.ShouldBindJSON(&param) != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.UserLogin(&param)
	if err != nil {
		response.ToErrorResponse(err)
	}
	response.ToResponse(user)

}
