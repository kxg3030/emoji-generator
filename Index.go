package main

import (
	"emoji/pkg/bootrap"
	"emoji/pkg/config"
	"emoji/pkg/controller/index"
	"github.com/gin-gonic/gin"
)

func main()  {
	DeleteExpireAssFile()
	bootstrap := bootrap.NewBootstrap(gin.New()).Init()
	bootstrap.Run(":8080")
}

func DeleteExpireAssFile()  {
	index.NewTask(config.RUNTIME_PATH).DeleteExpireAssFile()
}
