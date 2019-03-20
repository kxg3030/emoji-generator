package middleware

import (
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
				var msg string
				str,ok := err.(string)
				msg = str
				if ok == false{
					msg = ""
				}
				context.JSON(500,map[string]interface{}{
					"code"   : "500",
					"msg"    : msg,
					"status" : false,
					"data"   : map[string]string{},
				})
				return
			}
		}()
		context.Next()
	}
}
