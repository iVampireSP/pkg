package pkg

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

// JwtUserUrlToken 从 JWT 中获取用户信息。此函数只需要负责获取用户信息，不需要负责验证 JWT 的有效性。验证有效性在网关中完成。
func (GinMiddlewareStruct) JwtUserUrlToken(param string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var userInfo = ""

		// 从 Url 中获取
		var auth = c.Request.URL.Query().Get(param)

		if auth == "" {
			c.JSON(401, gin.H{
				"message": "缺少 JWT 令牌。请将 " + param + " 添加到查询中。",
			})
			c.Abort()
			return
		}

		// 接着取分割 payload
		payloadSplit := strings.Split(auth, ".")
		if len(payloadSplit) != 3 {
			c.JSON(401, gin.H{
				"message": "令牌格式不正确。",
			})
			c.Abort()
			return
		}

		// 取中间
		userInfo = payloadSplit[1]

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

		var structJwt = &JwtStruct{}
		err = sonic.Unmarshal(userInfoByte, structJwt)

		if err != nil {
			c.JSON(500, gin.H{
				"message": "我们无法传递您的信息，请稍后再试。" + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("auth", structJwt)
		c.Next()
	}
}
