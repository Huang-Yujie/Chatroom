package service

import (
	"github.com/Huang-Yujie/Chatroom/internal/model"
	"github.com/Huang-Yujie/Chatroom/internal/request"
	"github.com/Huang-Yujie/Chatroom/pkg/auth"
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
)

type LoginRespondContent struct {
	UserID   uint64 `json:"user_id"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

func (svc *Service) UserRegister(param *request.UserRegisterRequest) *errcode.Error {
	return svc.dao.UserRegister(param.UserName, param.Nickname, param.Password)
}

func (svc *Service) UserLogin(param *request.UserLoginRequest) (*LoginRespondContent, *errcode.Error) {
	user, err := svc.dao.UserLogin(param.UserName, param.Password)
	if err != nil {
		return nil, err
	}
	ID := user.ID
	token, _ := auth.GenerateToken(ID)
	return &LoginRespondContent{
		UserID:   ID,
		Nickname: user.Nickname,
		Token:    token,
	}, nil
}

func (svc *Service) UserGet(userID uint64) (*model.User, *errcode.Error) {
	return svc.dao.UserGet(userID)
}
