package pkg

import (
	"github.com/gin-gonic/gin"
)

// AuthUserKey 定义常量，以便在整个应用中使用这个常量来引用存储在 Context 中的键
const AuthUserKey = "auth"

// GetAuth 定义一个函数来从 gin.Context 中提取认证信息
func (UtilsStruct) GetAuth(c *gin.Context) (*JwtStruct, bool) {
	auth, exists := c.Get(AuthUserKey)
	if !exists {
		return nil, false
	}
	authUser, ok := auth.(*JwtStruct)
	return authUser, ok
}
