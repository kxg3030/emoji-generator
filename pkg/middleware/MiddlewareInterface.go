package middleware

import "github.com/gin-gonic/gin"

type MiddlewareInterface interface {
	Render() gin.HandlerFunc
}
