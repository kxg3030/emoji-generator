package middleware

import (
	"emoji/pkg/config"
	"emoji/pkg/system"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

type SslMiddleware struct {
	
}

func NewSslMiddleware() *SslMiddleware{
	return &SslMiddleware{

	}
}

func (this *SslMiddleware)Render()gin.HandlerFunc  {
	return func(context *gin.Context) {
		if context.Request.TLS != nil{
			context.Set("protocol","https://")
		}else{
			context.Set("protocol","http://")
		}
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost    : config.Config["ListenPort"].(string),
		})
		err := secureMiddleware.Process(context.Writer, context.Request)
		if err != nil {
			context.AbortWithStatusJSON(200,system.GetExceptionMessage(118))
			return
		}
		context.Next()
	}
}