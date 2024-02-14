package main

import (
	"github.com/ThirdWinter/Go/go-demo/route"
	"github.com/ThirdWinter/Go/mylog"
)

func main() {
	log.InitLog.Init(true, "Asia/Shanghai", "debug", "./", "log", ".log", 0644, 10*1024)
	r := route.Router()
	r.Run(":8080")
}
