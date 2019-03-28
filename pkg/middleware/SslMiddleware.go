package middleware

import (
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
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost    : "0.0.0.0:9527",
		})
		err := secureMiddleware.Process(context.Writer, context.Request)
		if err != nil {
			context.AbortWithStatusJSON(200,system.GetExceptionMessage(118))
			return
		}
		context.Next()
	}
}