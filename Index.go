package main

import (
	"emoji/pkg/bootstrap"
	"emoji/pkg/unity"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Print(unity.GetToken("openId","a"))
	engine := bootstrap.NewBootstrap(gin.New()).Init()
	engine.Run(":8080")
}

