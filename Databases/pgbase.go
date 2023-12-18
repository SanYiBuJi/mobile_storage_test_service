package Databases

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mobile_storage_test_service/logger"
	"sync"
	"time"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "123456"
	dbname   = "mobile-storage-test"
)

var (
	// DB 全局 db 模式
	//DB *gorm.DB
	// 单例工具
	dbOnce sync.Once
	// 连接 mysql 的 dsn
	dsn = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, host, port)
)

func Init() *gorm.DB {
	println(dsn)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Error(err.Error())
		panic("failed to connect database " + err.Error())
	}
	logger.Logger.Info("数据库连接创建成功～")
	err = SetConnect(DB)
	if err != nil {
		logger.Logger.Error("SetConnect failed : " + err.Error())
	}
	return DB
}

func SetConnect(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	return nil
}

func GetDBConnect() (*gorm.DB, error) {
	return Init(), nil
}
