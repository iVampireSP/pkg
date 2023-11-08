package pkg

import "github.com/gin-gonic/gin"

type utilsInterface interface {
	PasswordHash(pwd string) (string, error)
	PasswordVerify(pwd string, hash string) bool
}
type UtilsStruct struct{}

var Utils utilsInterface = UtilsStruct{}

type GinMiddlewareInterface interface {
	JwtUser() gin.HandlerFunc
	AllowAllCors() gin.HandlerFunc
}
type GinMiddlewareStruct struct{}

var GinMiddleware GinMiddlewareInterface = GinMiddlewareStruct{}
