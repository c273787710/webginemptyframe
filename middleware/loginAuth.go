package middleware

import (
	"adminframe/utils"
	"github.com/gin-gonic/gin"
	"adminframe/application/model"
	"adminframe/framework/config"
)

// #### jwt登录鉴权

func JWTMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		object := utils.NewObject(c)
		if token == "" {
			object.Response(utils.INVALID_AUTH_TOKEN,nil,"")
			c.Abort()
			return
		}
		info,err := utils.ParseJWTToken(token)
		if err != nil {
			object.Response(utils.INVALID_AUTH_TOKEN,nil,"")
			c.Abort()
			return
		}
		if config.AppSetting.AgentAuth == 1 && c.ClientIP() != info.ClientIP {
			object.Response(utils.INVALID_AUTH_TOKEN,nil,"")
			c.Abort()
			return
		}

		model,err := model.FindAdminByCondition(map[string]interface{}{"id =":info.ID})
		if err != nil || model == nil {
			object.Response(utils.INVALID_AUTH_TOKEN,nil,"")
			c.Abort()
			return
		}
		if model.Username != info.Username {
			object.Response(utils.INVALID_AUTH_TOKEN,nil,"")
			c.Abort()
			return
		}
		c.Set("uid",info.ID)
		c.Next()
	}
}