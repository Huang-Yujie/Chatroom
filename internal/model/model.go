package model

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/Huang-Yujie/Chatroom/internal/setting"
)

func NewDBEngine(databaseSettings *setting.DatabaseSetting) (*gorm.DB, error) {
	dsn := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn,
		databaseSettings.UserName,
		databaseSettings.Password,
		databaseSettings.Host,
		databaseSettings.DBName,
		databaseSettings.Charset,
		databaseSettings.ParseTime,
	)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(databaseSettings.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSettings.MaxOpenConns)
	return db, nil
}
