package middleware

import (
	"adminframe/utils"
	"github.com/gin-gonic/gin"
)

// #### jwt登录鉴权

func JWTMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token_valid := true
		if token == "" {
			token_valid = false
		}
		info,err := utils.ParseJWTToken(token)
		if err != nil {
			token_valid = false
		}
		if c.ClientIP() != info.ClientIP {
			token_valid = false
		}
		if !token_valid {
			object := utils.Object{C:c}
			object.Response(utils.INVALID_AUTH_TOKEN,nil,"")
			c.Abort()
			return
		}
		c.Next()
	}
}