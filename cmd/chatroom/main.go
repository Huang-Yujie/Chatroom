package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/Huang-Yujie/Chatroom/internal/routers"
	"github.com/Huang-Yujie/Chatroom/internal/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSettings()
	if err != nil {
		log.Fatalf("init.setupSettings err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSettings.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSettings.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSettings.ReadTimeout,
		WriteTimeout:   global.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe err: %v", err)
	}
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

	global.ServerSettings.ReadTimeout *= time.Second
	global.ServerSettings.WriteTimeout *= time.Second
	return nil
}
