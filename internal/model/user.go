package model

import (
	"github.com/Huang-Yujie/Chatroom/pkg/errcode"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `json:"user_id" gorm:"autoIncrement"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

func (u User) Create(db *gorm.DB) *errcode.Error {
	var user User
	err := db.Where("user_name = ?", u.UserName).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		return errcode.ErrorDuplicatedUserName
	}
	return errcode.Convert(db.Create(&user).Error)
}

func (u User) Get(db *gorm.DB) (User, *errcode.Error) {
	var user User
	err := db.Where("user_name = ?", u.UserName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, errcode.ErrorUserNameNotFound
	}
	return user, errcode.Convert(err)
}
