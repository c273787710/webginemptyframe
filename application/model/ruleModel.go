package model

import (
	"adminframe/framework/config"
	"time"
)

type RuleModel struct {
	ID int `gorm:"primary_key column:id" json:"id"`
	RuleName string `gorm:"column:rule_name" json:"rule_name"`
	RulePath string `gorm:"column:rule_path" json:"rule_path"`
	Method string `gorm:"column:method" json:"method"`
	Title string `gorm:"column:title" json:"title"`
	Pid int `gorm:"column:pid" json:"pid"`
	Auth int8 `gorm:"column:auth" json:"auth"`
	CreateTime int `gorm:"column:create_time" json:"create_time"`
	UpdateTime int `gorm:"column:update_time" json:"update_time"`
}

type RuleParam struct {
	ID int `json:"id"`
	RuleName string `json:"rule_name" binding:"required,max=30"`
	RulePath string `json:"rule_path" binding:"required,max=50"`
	Method string `json:"method" binding:"rulemethodvalidate"`
	Title string `json:"title" binding:"required,max=30"`
	Pid int `json:"pid" binding:"rulepidvalid"`
	Auth int8 `json:"auth" binding:"max=1"`
}


func (r RuleModel)TableName()string{
	return config.MysqlSetting.Prefix + "rule"
}

func FindRuleByCondition(condition map[string]interface{})(*RuleModel,error)  {
	whereSql,values,err := buildQuery(condition)
	if err != nil {
		return nil,err
	}
	model := new(RuleModel)
	err = DB.Where(whereSql,values...).Find(model).Error
	if err != nil {
		return nil,err
	}
	return model,err
}

func (r *RuleModel)UpdateRule(param RuleParam)error{
	r.RuleName = param.RuleName
	r.RulePath = param.RulePath
	if param.Method == "" {
		r.Method = "*"
	}else{
		r.Method = param.Method
	}
	r.Pid = param.Pid
	r.Auth = param.Auth
	r.Title = param.Title
	r.UpdateTime = int(time.Now().Unix())
	return DB.Save(r).Error
}

func CreateRule(param RuleParam)(*RuleModel,error){
	model := new(RuleModel)
	model.RuleName = param.RuleName
	model.RulePath = param.RulePath
	if param.Method == "" {
		model.Method = "*"
	}else{
		model.Method = param.Method
	}
	model.Title = param.Title
	model.Auth = param.Auth
	model.Pid = param.Pid
	model.CreateTime = int(time.Now().Unix())
	model.UpdateTime = int(time.Now().Unix())
	err := DB.Save(model).Error
	if err != nil{
		return nil,err
	}
	return model,err
}

func DeleteRule(id int)error{
	model := new(RuleModel)
	model.ID = id
	err := DB.First(model).Error
	if err != nil {
		return err
	}
	return DB.Delete(model).Error
}

func SelectRulesByCondition(condition map[string]interface{})([]RuleModel,error){
	var rules []RuleModel
	whereSql,values,err := buildQuery(condition)
	if err != nil {
		return rules,err
	}
	return rules,DB.Where(whereSql,values...).Find(&rules).Error
}