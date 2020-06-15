package admin

import (
	"github.com/gin-gonic/gin"
	"adminframe/utils"
	"adminframe/application/model"
	"strconv"
)

func AddRole(c *gin.Context){
	object := utils.NewObject(c)
	param := new(model.RoleParam)
	err := c.BindJSON(param)
	if err != nil {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	exit,_ := model.FindRoleByCondition(map[string]interface{}{"role_name":param.RoleName})
	if exit != nil {
		object.Response(utils.RECORD_IS_EXIT,nil,"")
		c.Abort()
		return
	}
	role,err := model.CreateRole(*param)
	if err != nil {
		object.Response(utils.CREATE_RECORD_ERR,nil,"")
		c.Abort()
		return
	}
	object.Response(utils.SUCCESS,role,"")
}

func UpdateRole(c *gin.Context){
	object := utils.NewObject(c)
	param := new(model.RoleParam)
	err := c.BindJSON(param)
	if err != nil {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	//判断id是否正确
	if param.ID == 0 || param.ID == param.Pid {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	//判断需要更新的是否存在
	role,err := model.FindRoleByCondition(map[string]interface{}{"id":param.ID})
	if err != nil || role == nil {
		object.Response(utils.RECORD_NOT_EXIT,nil,"")
		c.Abort()
		return
	}
	err = role.UpdateRole(*param)
	if err != nil {
		object.Response(utils.UPDATE_RECORD_ERR,nil,"")
		c.Abort()
		return
	}
	object.Response(utils.SUCCESS,role,"")
}

func DelRole(c *gin.Context){
	id , _ := strconv.Atoi(c.Query("id"))
	object := utils.NewObject(c)
	if id == 0 {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	err := model.DeleteRole(id)
	if err != nil {
		object.Response(utils.DELETE_RECORD_ERR,nil,"")
		c.Abort()
		return
	}
	object.Response(utils.SUCCESS,nil,"")
}

func ListRole(c *gin.Context){}
