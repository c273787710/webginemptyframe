package admin

import (
	"adminframe/application/model"
	"adminframe/framework/config"
	"adminframe/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"strings"
)
type adminParam struct {
	Username string `json:"username" binding:"required,min=5,max=30"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}

// #### 登录鉴权
func LoginAuth(c *gin.Context){
	object := utils.NewObject(c)
	var loginParam = new(adminParam)
	err := c.BindJSON(loginParam)
	if err != nil {
		//参数传递错误
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"用户名和密码错误")
		c.Abort()
		return
	}
	//判断用户名是否正确，并且判断是否超出登录错误次数
	admin,err := model.FindAdminByCondition(map[string]interface{}{"username":loginParam.Username})
	if err != nil {
		object.Response(utils.INVALID_AUTH_NAME,nil,"用户名不存在")
		c.Abort()
		return
	}

	if admin.LoginFailure >= int8(config.AppSetting.LoginFailureTime) &&
		int(time.Now().Unix()) - admin.UpdateTime < config.AppSetting.LoginFailureLock {
		object.Response(utils.ADMIN_LOCKING,nil,fmt.Sprintf("您失败次数超出限制，锁定半小时（剩余秒数:%d）",
			config.AppSetting.LoginFailureLock - (int(time.Now().Unix()) - admin.UpdateTime)))
		c.Abort()
		return
	}
	//执行登录
	admin.LastIP = c.ClientIP()
	err = admin.LoginAuth(loginParam.Password)
	if err != nil {
		object.Response(utils.ADMIN_AUTH_ERR,nil,err.Error())
		c.Abort()
		return
	}
	//登录成功，生成jwt鉴权码
	info := utils.UserTemplate{
		ID: admin.ID,
		Username: admin.Username,
		ClientIP: c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
	}
	token,err := utils.GenerateJWTToken(info)
	if err != nil {
		object.Response(utils.ADMIN_AUTH_ERR,nil,err.Error())
		c.Abort()
		return
	}
	object.Response(utils.SUCCESS,gin.H{
		"token":token,
		"user":admin,
	},"")
	c.Abort()
	return
}

// #### 获取用户信息
func AdminMenu(c *gin.Context){
	uid,_ := c.Get("uid")
	object := utils.NewObject(c)
	adminmodel,err := model.FindAdminByCondition(map[string]interface{}{"id":uid})
	if err != nil || adminmodel == nil {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	var rulenames []string
	if adminmodel.IsSup == 0 {
		rolemodel,_ := model.FindRoleByCondition(map[string]interface{}{"id":adminmodel.RoleID})
		query := map[string]interface{}{
			"id in":strings.Split(rolemodel.RuleIds,";"),
		}
		rulenames,err = model.GetRuleNameByRuleIDS(query)
	}
	object.Response(utils.SUCCESS,gin.H{
		"userinfo" : adminmodel,
		"rules":rulenames,
	},"")
}