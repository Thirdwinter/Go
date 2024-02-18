package core

import (
	"database/sql"
	"log"

	"github.com/ThirdWinter/Go/gvb_server/global"
	_ "github.com/go-sql-driver/mysql"
)

func SqlInit() (err error) {
	dsn := global.Config.Mysql.Dsn()
	global.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("dsn error: ", err)
	}
	err = global.DB.Ping()
	if err != nil {
		log.Panic("sql connect error: ", err)
	}
	global.DB.SetMaxOpenConns(100)
	global.DB.SetMaxIdleConns(10)
	log.Println("sql connect success")
	return nil
}


