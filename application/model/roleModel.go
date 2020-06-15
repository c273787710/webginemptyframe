package model

import (
	"adminframe/framework/config"
	"strings"
	"fmt"
	"time"
)

type RoleModel struct {
	ID int `gorm:"primary_key column:id" json:"id"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
	RuleIds string `gorm:"column:rule_ids" json:"rule_ids"`
	Pid int `gorm:"pid" json:"pid"`
	CreateTime int `gorm:"column:create_time" json:"create_time"`
	UpdateTime int `gorm:"column:update_time" json:"update_time"`
}

type RoleParam struct {
	ID int `json:"id"`
	RoleName string `json:"role_name" binding:"required,max=30"`
	RuleIds []int `json:"rule_ids" binding:"required,ruleidsisvalid"`
	Pid int `json:"pid" binding:"max=1,roleidisvalid"`
}

func (r RoleModel)TableName()string{
	return config.MysqlSetting.Prefix + "role"
}
func FindRoleByCondition(condition map[string]interface{})(*RoleModel,error){
	whereSql,values,err := buildQuery(condition)
	if err != nil {
		return nil,err
	}
	model := new(RoleModel)
	err = DB.Where(whereSql,values...).Find(model).Error
	if err != nil {
		return nil,err
	}
	return model,nil
}

func CreateRole(param RoleParam)(*RoleModel,error){
	model := new(RoleModel)
	model.RoleName = param.RoleName
	model.RuleIds = strings.Replace(strings.Trim(fmt.Sprint(param.RuleIds), "[]"), " ", ";", -1)
	model.Pid = param.Pid
	model.CreateTime = int(time.Now().Unix())
	model.UpdateTime = int(time.Now().Unix())
	err := DB.Save(model).Error
	if err != nil {
		return nil,err
	}
	return model,nil
}

func (r *RoleModel)UpdateRole(param RoleParam)error{
	r.RoleName = param.RoleName
	r.RuleIds = strings.Replace(strings.Trim(fmt.Sprint(param.RuleIds),"[]")," ",";",-1)
	r.Pid = param.Pid
	r.UpdateTime = int(time.Now().Unix())
	return DB.Save(r).Error
}

func DeleteRole(id int)error  {
	model := new(RoleModel)
	model.ID = id
	err := DB.First(model).Error
	if err != nil {
		return err
	}
	return DB.Delete(model).Error
}
