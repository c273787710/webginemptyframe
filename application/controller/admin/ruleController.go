package admin

import (
	"github.com/gin-gonic/gin"
	"adminframe/utils"
	"adminframe/application/model"
	"strconv"
)


func AddRule(c *gin.Context){
	param := new(model.RuleParam)
	err := c.BindJSON(param)
	object := utils.NewObject(c)
	if err != nil {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	query := map[string]interface{}{
		"rule_name": param.RuleName,
	}
	exit,_ := model.FindRuleByCondition(query)
	if exit != nil {
		object.Response(utils.RECORD_IS_EXIT,nil,"规则名称已存在")
		c.Abort()
		return
	}
	record,err := model.CreateRule(*param)
	if err != nil {
		object.Response(utils.CREATE_RECORD_ERR,nil,"创建规则错误")
		c.Abort()
		return
	}
	object.Response(utils.SUCCESS,record,"")
}

func UpdateRule(c *gin.Context){
	object := utils.NewObject(c)
	param := new(model.RuleParam)
	err := c.BindJSON(param)
	if err != nil {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	if param.ID == 0 || param.ID == param.Pid {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"规则ID错误")
		c.Abort()
		return
	}
	rule,_ := model.FindRuleByCondition(map[string]interface{}{"id":param.ID})
	if rule == nil{
		object.Response(utils.RECORD_NOT_EXIT,nil,"")
		c.Abort()
		return
	}
	//判断rulename是否存在
	exit,_ := model.FindRuleByCondition(map[string]interface{}{"id !=":param.ID,"rule_name":param.RuleName})
	if exit != nil {
		object.Response(utils.RECORD_IS_EXIT,nil,"rule_name已存在")
		c.Abort()
		return
	}
	err = rule.UpdateRule(*param)
	if err != nil {
		object.Response(utils.UPDATE_RECORD_ERR,nil,"")
		c.Abort()
		return
	}
	object.Response(utils.SUCCESS,rule,"")
}

func DelRule(c *gin.Context){
	idstring := c.Query("id")
	id,_ := strconv.Atoi(idstring)
	object := utils.NewObject(c)
	if id == 0 {
		object.Response(utils.INVALID_REQUEST_PARAMS,nil,"")
		c.Abort()
		return
	}
	err := model.DeleteRule(id)
	if err != nil {
		object.Response(utils.DELETE_RECORD_ERR,nil,"")
		c.Abort()
		return
	}
	object.Response(utils.SUCCESS,nil,"")
}

func ListRule(c *gin.Context){}
