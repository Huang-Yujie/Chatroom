package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Huang-Yujie/Chatroom/global"
	"github.com/Huang-Yujie/Chatroom/internal/model"
	"github.com/Huang-Yujie/Chatroom/internal/routers"
	"github.com/Huang-Yujie/Chatroom/internal/setting"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	err := setupSettings()
	if err != nil {
		log.Fatalf("init.setupSettings err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = SetupTableModel(global.DBEngine, &model.User{}, &model.Message{})
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

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSettings)
	if err != nil {
		return err
	}
	return nil
}

func SetupTableModel(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models)
	if err != nil {
		return err
	}
	return nil
}
