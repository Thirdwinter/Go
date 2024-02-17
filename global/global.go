package global

import (
	"github.com/ThirdWinter/Go/gvb_server/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB *gorm.DB
)