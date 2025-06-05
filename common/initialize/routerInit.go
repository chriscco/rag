package initialize 

import (
	"github.com/gin-gonic/gin"
	"main/router"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	allRouters := router.AllRouters 
	base := r.Group("/")
	{
		allRouters.RagRouter.ApiRouterInit(base)
	}
	return r 
}