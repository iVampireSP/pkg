package pkg

import (
	"encoding/base64"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

// JwtUser 从 JWT 中获取用户信息。此函数只需要负责获取用户信息，不需要负责验证 JWT 的有效性。验证有效性在网关中完成。
func (GinMiddleware) JwtUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := c.Request.Header.Get("X-Jwt-Payload")

		// 检测有无 X-Jwt-Payload
		if userInfo == "" {
			c.JSON(401, gin.H{
				"message": "无法认证用户。",
			})
			c.Abort()
			return
		}

		// base64 decode
		userInfoByte, err := base64.RawURLEncoding.DecodeString(userInfo)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "获取用户信息时发生错误，可能是网关发生了问题。" + err.Error(),
			})
			c.Abort()
			return
		}

		var jwtUser = map[string]interface{}{}

		err = sonic.Unmarshal(userInfoByte, &jwtUser)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "我们无法传递您的信息，请稍后再试。" + err.Error(),
			})
			c.Abort()
			return
		}

		for k, v := range jwtUser {
			c.Set("auth."+k, v)
		}

		// 在之后，获取用户信息就用 c.Get("auth.user") 即可。
		c.Next()
	}
}
