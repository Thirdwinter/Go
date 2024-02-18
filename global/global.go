package global

import (
	"database/sql"

	"github.com/ThirdWinter/Go/gvb_server/config"
)

var (
	Config *config.Config
	DB     *sql.DB
)
