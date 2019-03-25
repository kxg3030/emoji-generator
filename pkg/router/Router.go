package router

import (
	"emoji/pkg/controller/admin"
	"emoji/pkg/controller/index"
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
	{
		groupIndex.GET("/user"      ,index.NewEmoji().EmojiGenerator)
		groupIndex.GET("/emoji/list",index.NewEmojiFile().GetEmojiFileList)
		groupIndex.GET("/emoji/login",index.NewUserList().Login)
	}
	groupAdmin := this.Router.Group("/api/emoji")
	{
		groupAdmin.POST("/upload"        ,admin.NewEmoji().UploadFile)
		groupAdmin.GET("/generator/:code",admin.NewEmoji().GeneratorGifFromVideo)
	}
	return  this.Router
}
