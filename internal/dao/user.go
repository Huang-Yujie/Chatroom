package dao

import (
	"github.com/Huang-Yujie/Chatroom/internal/model"
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
)

func (d *Dao) UserRegister(username, nickname, password string) (*model.User, *errcode.Error) {
	user := model.User{UserName: username, Nickname: nickname, Password: password}
	err := user.Create(d.engine)
	if err != nil {
		return nil, err
	}
	return d.UserLogin(username, password)
}

func (d *Dao) UserLogin(username, password string) (*model.User, *errcode.Error) {
	user, err := model.User{UserName: username}.Get(d.engine)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errcode.ErrorPassword
	}
	return &model.User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
