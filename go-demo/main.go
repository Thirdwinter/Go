package main

import (
	"github.com/ThirdWinter/Go/go-demo/route"
	log "github.com/ThirdWinter/Go/mylog"
)

func main() {
	log.Init(true,  "debug", "./", "log.log", 10*1024)
	r := route.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
