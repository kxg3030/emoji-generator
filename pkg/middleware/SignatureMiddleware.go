package middleware

import (
	"emoji/pkg/system"
	"emoji/pkg/unity"
	"github.com/gin-gonic/gin"
	"html"
)

const key  = "emoji"
type SignatureMiddleware struct {

}

func NewSignatureMiddleware()*SignatureMiddleware  {
	return &SignatureMiddleware{

	}
}

func (this *SignatureMiddleware)Render()gin.HandlerFunc  {
	return func(context *gin.Context) {
		timestamp := html.EscapeString(context.Query("timestamp"))
		signature := html.EscapeString(context.Query("signature"))
		if unity.Md5String(timestamp + key) == signature{
			context.Next()
		}
		context.AbortWithStatusJSON(200,system.GetExceptionMessage(111))
	}
}