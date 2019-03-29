package router

import (
	"emoji/pkg/config"
	"emoji/pkg/controller/admin"
	"emoji/pkg/controller/index"
	"emoji/pkg/middleware"
	"emoji/pkg/system"
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
	this.Router.NoRoute(this.RegisterNotFound())
	groupIndex := this.Router.Group("/api/v1")
	// 用户登陆
	groupIndex.GET("/user/login",index.NewUserList().Login)
	this.RegisterIndexMiddleWare(groupIndex)
	{
		// 用户创建GIF
		groupIndex.POST("/emoji/create",index.NewUserEmojiFile().EmojiGenerator)
		// 获取封面列表
		groupIndex.GET("/emoji/list" ,index.NewSysEmojiFile().GetEmojiFileList)
	}

	groupAdmin := this.Router.Group("/api/emoji")
	this.RegisterAdminMiddleWare(groupAdmin)
	{
		// 后台上传文件
		groupAdmin.POST("/upload"        ,admin.NewSysEmojiFile().UploadFile)
		// 生成封面图和GIF
		groupAdmin.GET("/generator/:encode",admin.NewSysEmojiFile().GeneratorGifFromVideo)
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

func (this *Router)RegisterNotFound()gin.HandlerFunc  {
	return func(context *gin.Context) {
		context.AbortWithStatusJSON(404,system.GetExceptionMessage(119))
	}
}
