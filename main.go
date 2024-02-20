package main

import (
	"github.com/ThirdWinter/Go/gvb_server/core"
	"github.com/ThirdWinter/Go/gvb_server/routers"
	log "github.com/ThirdWinter/Go/mylog"
)

func main() {
	core.InitConf()
	core.Initlog()
	// 读取配置文件
	// err:=core.SqlInit()
	// if err != nil {
	// 	panic(err)
	// }
	// test.InsertRowDemo()
	core.InitGorm()
	routers.InitRouter()
	log.Debug("这是一次启动调试")
	//global.DB = core.InitGorm()
	//fmt.Println(global.DB)
}
