package main

import (
	"github.com/ThirdWinter/Go/gvb_server/core"
	"github.com/ThirdWinter/Go/gvb_server/global"
	"github.com/ThirdWinter/Go/gvb_server/test"
	log "github.com/ThirdWinter/Go/mylog"
)

func main() {
	core.InitConf()
	// 读取配置文件
	//size,err:=strconv.Atoi(global.Config.Logger.Size)

	//fmt.Println(err,size)

	log.Init(true, "debug", "./", "log.log", global.Config.Logger.Size)
	log.Debug("111")
	err:=core.SqlInit()
	if err != nil {
		panic(err)
	}
	test.InsertRowDemo()
	//global.DB = core.InitGorm()
	//fmt.Println(global.DB)
}
