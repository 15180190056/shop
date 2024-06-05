package dbserver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mystopproject/conf"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var dsn string = "root:123456@tcp(121.40.117.237:3306)/SpotProject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	// 自动迁移模式，根据模型创建表
	// 注意：在生产环境中，通常使用迁移工具来管理数据库结构
	db.AutoMigrate(&conf.User{}, &conf.Product{}, &conf.Order{})

	return db, nil
}
