package main

import (
	"emoji/pkg/bootstrap"
	"emoji/pkg/config"
	"emoji/pkg/unity"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Println(unity.GetToken("openId","x"))
	setRunMode()
	engine := bootstrap.NewBootstrap(gin.New()).Init()
	engine.Run(true)
}

func setRunMode()  {
	if config.Config["DebugMode"].(bool){
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}
}

