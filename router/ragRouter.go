package router

import (
	"main/api/controller"
	"main/api/service"

	"github.com/gin-gonic/gin"
)

type RagRouter struct {
	service service.RagServiceImpl 
} 
func (rr *RagRouter) ApiRouterInit(router *gin.RouterGroup) {
	r := router.Group("/rag")
	rr.service = service.NewRagService() 
	ragCtrl := controller.NewRagController(rr.service) 
	{
		r.GET("/", ragCtrl.Index)
		r.POST("/query", ragCtrl.Query)
		r.POST("/upload", ragCtrl.Upload)
	}
}
