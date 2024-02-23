package main

import (
	"github.com/ThirdWinter/Go/gvb_server/core"
	"github.com/ThirdWinter/Go/gvb_server/routers"
)

func main() {
	core.InitConf()
	core.Initlog()
	core.InitGorm()
	routers.InitRouter()
}
