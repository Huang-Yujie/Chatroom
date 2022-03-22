package model

import (
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `json:"user_id" gorm:"autoIncrement"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"-"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (u User) Create(db *gorm.DB) *errcode.Error {
	var user User
	err := db.Where("user_name = ?", u.UserName).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		return errcode.ErrorDuplicatedUserName
	}
	err = db.Create(&u).Error
	if err != nil {
		return errcode.Convert(err)
	}
	return nil
}

func (u User) Get(db *gorm.DB) (User, *errcode.Error) {
	var user User
	err := db.Where("user_name = ?", u.UserName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, errcode.ErrorUserNameNotFound
	}
	if err != nil {
		return user, errcode.Convert(err)
	}
	return user, nil
}
