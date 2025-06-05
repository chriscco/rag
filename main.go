package main

import (
	"main/common/initialize"
	"main/global"
)

func main() {
	router := initialize.GlobalInit() 
	// TODO: Set router runnning level 
	router.Run(":8080")
	global.Log.Info("server booted")
}