package main

import (
	"emoji/pkg/bootrap"
	"github.com/gin-gonic/gin"
)

func main()  {
	bootstrap := bootrap.NewBootstrap(gin.New()).Init()
	bootstrap.Run(":8080")
}

