package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/Huang-Yujie/Chatroom/internal/setting"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	setupSettings()
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://175.178.43.145:4001/ws?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJpc3MiOiJIdWFuZyIsImV4cCI6MTY0ODEwNDEzMSwibmJmIjoxNjQ4MDk2OTMxLCJpYXQiOjE2NDgwOTY5MzF9.pvLgLeNv_jaP9nu7HV03CFcm-tGrsPwi5q9d1HyR0U4", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "内部错误！")

	err = wsjson.Write(ctx, c, "Hello WebSocket Server")
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("接收到服务端响应：%v\n", v)

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
