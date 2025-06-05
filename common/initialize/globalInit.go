package initialize

import (
	"main/common/config"
	"main/common/log"
	"main/global"
	"github.com/gin-gonic/gin"
)

func GlobalInit() *gin.Engine {
	global.Config = config.ConfigInit() 
	global.Log = logger.NewLogger(global.Config.Log.FilePath) 
	router := RouterInit() 
	return router
}