package global

import (
	//"database/sql"

	//"github.com/ThirdWinter/Go/gvb_server/config"
	"github.com/ThirdWinter/Go/gvb_server/config"
	"github.com/jinzhu/gorm"
)

var (
	Config *config.Config
	//DB     *sql.DB
	Db *gorm.DB
)
