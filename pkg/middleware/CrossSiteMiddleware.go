package middleware

import "github.com/gin-gonic/gin"

type CrossSiteMiddleware struct {

}

func NewCrossSiteMiddleware()*CrossSiteMiddleware  {
	return &CrossSiteMiddleware{

	}
}

func (this *CrossSiteMiddleware)Render()gin.HandlerFunc  {
	return func(context *gin.Context) {
		headers:= context.Request.Header.Get("Access-Control-Request-Headers")
		origin := context.Request.Header.Get("origin")
		method := context.Request.Method
		if origin == "http://127.0.0.1:8080" || origin == "http://127.0.0.1:8081" {
			// 允许请求的域
			context.Header("Access-Control-Allow-Origin", origin)
			// 允许请求的header头
			context.Header("Access-Control-Allow-Headers", headers)
			// 允许请求的方法类型
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			// 允许请求的缓存时间
			context.Header("Access-Control-Max-Age", "600")
			// 是否验证cookie
			context.Header("Access-Control-Allow-Credentials", "true")
			// 返回的数据内容是否缓存
			context.Header("Cache-Control", "no-store")
			// 返回的数据格式
			context.Set("Content-Type", "application/json")
		}
		if method == "options"{
			context.AbortWithStatus(204)
		}else{
			context.Next()
		}
	}
}
