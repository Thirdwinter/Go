package dao

import (
	"time"

	"github.com/ThirdWinter/Go/go-demo/config"
	log "github.com/ThirdWinter/Go/mylog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		log.Error("mysql connect error:%s", err)
		return
	}
	if Db.Error != nil {
		log.Error("db error:%s", Db.Error)
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
}
