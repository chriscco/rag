package controller

import (
	"main/api/entity"
	"main/api/service"
	"main/common/retcode"
	"main/global"

	"github.com/gin-gonic/gin"
)

// RagController
type RagController struct {
	ragService service.RagServiceImpl
}

// NewRagController 
//	@return RagController 
func NewRagController(service service.RagServiceImpl) *RagController {
	return &RagController{
		ragService: service,
	}
}

// Query Send user input to RAG model API and retrieves result  
//	@param c 
func (rc *RagController) Query(c *gin.Context) {
	var request entity.Request 
	err := c.Bind(&request)
	if err != nil {
		global.Log.Warn(err, "unable to bind")
		retcode.Fatal(c, err, "")
		return
	}
	global.Log.Info("query request: ", request.Query)
	resp, err := rc.ragService.CallAPI() 
	if err != nil {
		global.Log.Fatal("unable to call CallAPI ", err) 
		retcode.Fatal(c, err, "")
	}
	response := entity.Response {
		Result: resp,
	}
	retcode.Ok(c, response)
}

func (rc *RagController) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H {})
}

// Upload 
//	@param c 
func (rc *RagController) Upload(c *gin.Context) {
	
}
