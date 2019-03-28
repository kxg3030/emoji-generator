package main

import (
	"emoji/pkg/bootstrap"
	"emoji/pkg/config"
	"github.com/gin-gonic/gin"
)

func main()  {
	setRunMode()
	engine := bootstrap.NewBootstrap(gin.New()).Init()
	engine.Run(":9527")
}

func setRunMode()  {
	if config.Config["DebugMode"].(bool){
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}
}

