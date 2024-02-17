package main

import (
	"fmt"

	"github.com/ThirdWinter/Go/gvb_server/core"
	"github.com/ThirdWinter/Go/gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
}