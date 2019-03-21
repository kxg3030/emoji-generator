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
	emojiIndex := index.NewEmoji()
	emojiAdmin := admin.NewEmoji()
	groupIndex := this.Router.Group("/api/v1")
	{
		groupIndex.GET("/user", emojiIndex.EmojiGenerator)
	}
	groupAdmin := this.Router.Group("/api/emoji")
	{
		groupAdmin.POST("/upload"   ,emojiAdmin.UploadFile)
		groupAdmin.GET("/generator/:code",emojiAdmin.GeneratorGifFromVideo)
	}
	return  this.Router
}
