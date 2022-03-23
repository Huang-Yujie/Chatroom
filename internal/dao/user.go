package dao

import (
	"github.com/Huang-Yujie/Chatroom/internal/model"
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
)

func (d *Dao) UserRegister(username, nickname, password string) *errcode.Error {
	user := model.User{UserName: username, Nickname: nickname, Password: password}
	return user.Create(d.engine)
}

func (d *Dao) UserLogin(username, password string) (*model.User, *errcode.Error) {
	user, err := model.User{ID: 0, UserName: username}.Get(d.engine)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errcode.ErrorPassword
	}
	return user, nil
}

func (d *Dao) UserGet(userID uint64) (*model.User, *errcode.Error) {
	return model.User{ID: userID}.Get(d.engine)
}
