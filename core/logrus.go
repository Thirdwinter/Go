package core

import (
	"github.com/ThirdWinter/Go/gvb_server/global"
	log "github.com/ThirdWinter/Go/mylog"
)

func Initlog() {
	log.Init(global.Config.Logger.Out, global.Config.Logger.Level, global.Config.Logger.Path, global.Config.Logger.Name, global.Config.Logger.Size)
}