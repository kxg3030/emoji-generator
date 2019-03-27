package middleware

import (
	"emoji/pkg/logger"
	"emoji/pkg/system"
	"github.com/gin-gonic/gin"
)

type RecoverMiddleware struct {

}

func NewRecoverMiddleware() *RecoverMiddleware  {
	return &RecoverMiddleware{

	}
}

func (this *RecoverMiddleware)Render()gin.HandlerFunc  {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Error(err.(string))
				context.JSON(500,system.GetExceptionMessage(108))
				return
			}
		}()
		context.Next()
	}
}
