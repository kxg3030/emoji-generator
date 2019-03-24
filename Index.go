package main

import (
	"emoji/pkg/bootstrap"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := bootstrap.NewBootstrap(gin.New()).Init()
	engine.Run(":8080")
}

