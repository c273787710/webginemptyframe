package middleware

import (
	"adminframe/application/model"
	"adminframe/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

//权限验证中间件

func AdminAuthMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
		adminid := c.GetInt("uid")
		object := utils.NewObject(c)
		if adminid == 0 {
			object.Response(utils.INVALID_PERMISSION_OPERATE,nil,"")
			c.Abort()
			return
		}
		rule := strings.Trim(c.Request.URL.Path,"/")
		method := strings.ToLower(c.Request.Method)
		rulemodel,_ := model.FindRuleByCondition(map[string]interface{}{"rule_path":rule,"method":method})
		if rulemodel == nil || rulemodel.Auth == 0{
			c.Next()
			return
		}
		//需要权限
		admininfo,_ := model.FindAdminByCondition(map[string]interface{}{"id":adminid})
		if admininfo == nil {
			object.Response(utils.INVALID_AUTH_TOKEN,nil,"")
			c.Abort()
			return
		}
		if admininfo.IsSup == 1 {
			c.Next()
			return
		}
		if  admininfo.RoleID == 0 {
			object.Response(utils.INVALID_PERMISSION_OPERATE,nil,"")
			c.Abort()
			return
		}
		rolemodel,_ := model.FindRoleByCondition(map[string]interface{}{"id":admininfo.RoleID})
		if rolemodel == nil {
			object.Response(utils.INVALID_PERMISSION_OPERATE,nil,"")
			c.Abort()
			return
		}
		rules := strings.Split(rolemodel.RuleIds,";")
		for _,v := range rules{
			if v == strconv.Itoa(rulemodel.ID) {
				c.Next()
				return
			}
		}
		object.Response(utils.INVALID_PERMISSION_OPERATE,nil,"")
		c.Abort()
		return
	}
}
