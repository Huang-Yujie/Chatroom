package main

import (
	"fmt"
	"time"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/Huang-Yujie/Chatroom/internal/setting"
	"github.com/Huang-Yujie/Chatroom/pkg/auth"
)

func main() {
	setupSettings()
	// token, err := auth.GenerateToken(4)
	// if err != nil {
	// 	panic(err)
	// }
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJpc3MiOiJIdWFuZyIsImV4cCI6MTY0ODAyODI2NCwibmJmIjoxNjQ4MDI4MjY0LCJpYXQiOjE2NDgwMjgyNjR9.Hn7VBIEOi02Hhe7uCxD4AoMzeCckizQ9SjgIL7EiUD4"
	claims, err := auth.ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println("UserID", claims.UserID)
}
func setupSettings() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settings.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	}
	err = settings.ReadSection("Database", &global.DatabaseSettings)
	if err != nil {
		return err
	}
	err = settings.ReadSection("JWT", &global.JWTSettings)
	if err != nil {
		return err
	}

	global.ServerSettings.ReadTimeout *= time.Second
	global.ServerSettings.WriteTimeout *= time.Second
	global.JWTSettings.Expire *= time.Second
	return nil
}
