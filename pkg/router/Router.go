package router

import (
	"emoji/pkg/config"
	"emoji/pkg/controller/admin"
	"emoji/pkg/controller/index"
	"emoji/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router *gin.Engine
}

func NewRouter(frameworkRouter *gin.Engine)*Router  {
	return &Router{
		Router:frameworkRouter,
	}
}

// register router
func (this *Router)RegisterRouter()*gin.Engine  {

	groupIndex := this.Router.Group("/api/v1")
	groupIndex.GET("/user/login",index.NewUserList().Login)
	this.RegisterIndexMiddleWare(groupIndex)
	{
		groupIndex.GET("/user"       ,index.NewEmoji().EmojiGenerator)
		groupIndex.GET("/emoji/list" ,index.NewEmojiFile().GetEmojiFileList)
	}

	groupAdmin := this.Router.Group("/api/emoji")
	{
		groupAdmin.POST("/upload"        ,admin.NewEmoji().UploadFile)
		groupAdmin.GET("/generator/:code",admin.NewEmoji().GeneratorGifFromVideo)
	}
	return  this.Router
}

func (this *Router)RegisterIndexMiddleWare(group *gin.RouterGroup)*gin.RouterGroup  {
	middlewareHandle,ok := config.Config["LocalMiddleWare"].(map[string][]middleware.MiddlewareInterface)
	if ok{
		for _,val := range middlewareHandle["index"]{
			group.Use(val.Render())
		}
	}
	return group
}

func (this *Router)RegisterAdminMiddleWare(group *gin.RouterGroup)*gin.RouterGroup  {
	middlewareHandle,ok := config.Config["LocalMiddleWare"].(map[string][]middleware.MiddlewareInterface)
	if ok{
		for _,val := range middlewareHandle["admin"]{
			group.Use(val.Render())
		}
	}
	return group
}
