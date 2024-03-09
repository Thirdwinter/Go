package core

import (
	"time"

	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/models"
	//"github.com/ThirdWinter/Go/gvb_server/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// import (
// 	"database/sql"
// 	"log"

// 	"github.com/ThirdWinter/Go/gvb_server/global"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func SqlInit() (err error) {
// 	dsn := global.Config.Mysql.Dsn()
// 	global.DB, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Panic("dsn error: ", err)
// 	}
// 	err = global.DB.Ping()
// 	if err != nil {
// 		log.Panic("sql connect error: ", err)
// 	}
// 	global.DB.SetMaxOpenConns(100)
// 	global.DB.SetMaxIdleConns(10)
// 	log.Println("sql connect success")
// 	return nil
// }

var err error

func InitGorm() {
	global.Db, err = gorm.Open("mysql", global.Config.Mysql.Dsn())
	if err != nil {
		panic(err)
	}
	global.Db.DB().SetMaxIdleConns(10)
	global.Db.DB().SetMaxOpenConns(100)
	global.Db.DB().SetConnMaxLifetime(10 * time.Second)
	global.Db.SingularTable(true)                                                // 禁止表名复数化
	global.Db.AutoMigrate(&models.User{}, &models.Category{}, &models.Article{}) // 数据模型迁移
	//global.Db.Close()
}
