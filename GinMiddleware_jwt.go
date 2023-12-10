package pkg

import (
	"encoding/base64"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

var appDebug = os.Getenv("DEBUG") == "true"

type JwtStruct struct {
	Iss    string      `json:"iss"`
	Iat    int         `json:"iat"`
	Exp    int         `json:"exp"`
	Nbf    int         `json:"nbf"`
	Jti    string      `json:"jti"`
	Sub    string      `json:"sub"`
	Prv    string      `json:"prv"`
	TeamId interface{} `json:"team_id"`
	User   struct {
		Id                 uint        `json:"id"`
		Uuid               string      `json:"uuid"`
		Name               string      `json:"name"`
		Email              string      `json:"email"`
		EmailVerifiedAt    time.Time   `json:"email_verified_at"`
		RealNameVerifiedAt interface{} `json:"real_name_verified_at"`
	} `json:"user"`
}

// JwtUser 从 JWT 中获取用户信息。此函数只需要负责获取用户信息，不需要负责验证 JWT 的有效性。验证有效性在网关中完成。
func (GinMiddlewareStruct) JwtUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userInfo = ""
		// 从 Authorization 中获取用户信息
		// 分割 bearer
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{
				"message": "缺少 JWT 令牌。",
			})
			c.Abort()
			return
		}

		authSplit := strings.Split(auth, " ")
		if len(authSplit) != 2 {
			c.JSON(401, gin.H{
				"message": "无效的 JWT 令牌。",
			})
			c.Abort()
			return
		}

		// 接着取分割 payload
		payloadSplit := strings.Split(authSplit[1], ".")
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

		//  var h, _ = c.Get("auth")
		//
		//	user := h.(*utils.JwtStruct)
		//	fmt.Println(user)
		//	return
		c.Next()
	}
}
