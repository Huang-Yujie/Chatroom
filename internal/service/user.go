package service

import (
	"github.com/Huang-Yujie/Chatroom/internal/model"
	"github.com/Huang-Yujie/Chatroom/internal/request"
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
)

func (svc *Service) UserRegister(param *request.UserRegisterRequest) (*model.User, *errcode.Error) {
	return svc.dao.UserRegister(param.UserName, param.Nickname, param.Password)
}

func (svc *Service) UserLogin(param *request.UserLoginRequest) (*model.User, *errcode.Error) {
	return svc.dao.UserLogin(param.UserName, param.Password)
}
