package middleware

import (
	"emoji/pkg/system"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
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
		token,err := request.ParseFromRequest(context.Request,request.AuthorizationHeaderExtractor, func(token *jwt.Token) (i interface{}, err error) {
			return []byte("emoji"),nil
		})
		if err == nil {
			if token.Valid {
				if val,ok:= token.Claims.(jwt.MapClaims);ok{
					context.Set("openId",val["openId"])
				}
				context.Next()
			}else {
				context.AbortWithStatusJSON(200,system.GetExceptionMessage(106))
			}
		}else{
			context.AbortWithStatusJSON(200,system.GetExceptionMessage(107))
		}
	}
}