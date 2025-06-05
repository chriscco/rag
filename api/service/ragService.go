package service

import (
	"fmt"
	"io"
	"main/global"
	"net/http"
)

type RagServiceImpl struct {}

func NewRagService() RagServiceImpl {
	return RagServiceImpl{}
}

func (rs *RagServiceImpl) CallAPI() (string, error){
	api, host, port := global.Config.Api.ApiPath, 
					   global.Config.Api.ApiHost, 
					   global.Config.Api.ApiPort 
	return rs.requestAPI(api, host, port) 
}

func (rs *RagServiceImpl) requestAPI(api, host, port string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/%s", host, port, api)) 
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