package global

import (
	"main/common/log"
	"main/common/config"
)

var (
	Log    logger.ILog
	Config *config.AllConfig 
)