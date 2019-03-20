package bootrap

import (
	"emoji/pkg/config"
	"emoji/pkg/database"
	"emoji/pkg/middleware"
	"emoji/pkg/router"
	"emoji/pkg/unity"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
	"net/http"
)


type Bootstrap struct {
	Framework  *gin.Engine
}


func NewBootstrap(framework *gin.Engine)*Bootstrap  {
	return &Bootstrap{
		Framework:framework,
	}
}

func (this *Bootstrap)Init()*Bootstrap  {
	this.setDebugMode()
	this.setMiddleware()
	this.setAssetsPath()
	this.initFrameworkRouter()
	this.setOrm()
	return this
}

func (this *Bootstrap)Run(port string)  {
	unity.ErrorCheck(this.Framework.Run(port))
}

func (this *Bootstrap)initFrameworkRouter()*Bootstrap  {
	this.Framework = router.NewRouter(this.Framework).RegisterRouter()
	return this
}

func (this *Bootstrap)setDebugMode()  {
	debugMode := config.Config["DebugMode"].(bool)
	if debugMode{
		this.Framework.Use(gin.Logger())
	}
}

func (this *Bootstrap)setMiddleware()  {
	middlewareHandle,ok := config.Config["MiddleWare"].([]middleware.MiddlewareInterface)
	if ok{
		for _,val := range middlewareHandle{
			this.Framework.Use(val.Render())
		}
	}
}

func (this *Bootstrap)setAssetsPath()  {
	this.Framework.StaticFS("/assets",http.Dir(config.ASSETS_PATH))
}

func (this *Bootstrap)setOrm()  {
	var err error
	database.Database,err = gorose.Open(config.Database)
	unity.ErrorCheck(err)
}



