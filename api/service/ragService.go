package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/global"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type RagServiceImpl struct {
	imgPath string
}

func NewRagService() RagServiceImpl {
	return RagServiceImpl{
		imgPath: "resource/image",
	}
}

func (rs *RagServiceImpl) CallAPI(c *gin.Context, message string) (string, error){
	api, host, port := global.Config.Api.ApiPath, 
					   global.Config.Api.ApiHost, 
					   global.Config.Api.ApiPort 
	return rs.requestAPI(api, host, port, message) 
}

func (rs *RagServiceImpl) requestAPI(api, host, port string, message string) (string, error) {
	request := map[string]interface{} {
		"message": message, 
		"imgPath": rs.imgPath, 
	}
	jsonBytes, _ := json.Marshal(request)
	reqBody := bytes.NewBuffer(jsonBytes) 

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%s/%s", host, port, api), 
		"application/json",
		reqBody) 
	if err != nil {
		global.Log.Fatal("unable to call API ", err) 
		return "", nil 
	}
	defer resp.Body.Close() 
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Log.Fatal("unable to read response body ", err) 
		return "", nil 
	}
	return string(body), nil 
}

func (rs *RagServiceImpl) SaveImg(c *gin.Context) (error) {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	if _, err := os.Stat(rs.imgPath); os.IsNotExist(err) {
		os.MkdirAll(rs.imgPath, os.ModePerm)
	}

	filename := fmt.Sprintf("%d.jpg", time.Now().Unix())
	dst := filepath.Join(rs.imgPath, filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return err 
	}
	return nil 
}