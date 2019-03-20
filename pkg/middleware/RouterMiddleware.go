package middleware

import (
	"github.com/gin-gonic/gin"
)

type RouterMiddleware struct {

}

func NewRouterMiddleware()*RouterMiddleware  {
	return &RouterMiddleware{

	}
}

func (this *RouterMiddleware)Render() gin.HandlerFunc  {
	return func(context *gin.Context) {
		context.Next()
	}
}