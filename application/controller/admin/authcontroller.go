package admin

import (
	"adminframe/application/model"
	"adminframe/framework/config"
	"adminframe/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
type adminParam struct {
	Username string `json:"username" binding:"required,min=5,max=30"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}

// #### 登录鉴权
func LoginAuth(c *gin.Context){
	object := utils.Object{C:c}
	var loginParam = new(adminParam)
	err := c.BindJSON(loginParam)
	if err != nil {
		//参数传递错误
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,err.Error())
		c.Abort()
		return
	}
	//判断用户名是否正确，并且判断是否超出登录错误次数
	admin,err := model.FindAdminByCondition(map[string]interface{}{"username":loginParam.Username})
	if err != nil {
		object.Response(utils.INVALID_AUTH_NAME,nil,err.Error())
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
	object.Response(utils.SUCCESS,gin.H{"token":token},"")
	c.Abort()
	return
}

// #### 获取用户信息
func AdminInfo(c *gin.Context){

}