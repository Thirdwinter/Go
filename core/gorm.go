package core

import (
	"log"
	"time"
	"fmt"
	"github.com/ThirdWinter/Go/gvb_server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == ""{
		log.Println("未配置mysql, 取消连接")
		return nil
	}
	dsn:=global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env=="dev"{
		// 开发环境
		mysqlLogger =logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger =logger.Default.LogMode(logger.Error)
	}

	db,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}

	sqlDB,_:=db.DB()
	sqlDB.SetMaxIdleConns(10) // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100) // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour*4) // 连接最大复用时间
	return db
}
