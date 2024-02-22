package global

import (
	"github.com/ThirdWinter/Go/gvb_server/config"
	"github.com/jinzhu/gorm"
)

var (
	Config *config.Config
	Db *gorm.DB
)
