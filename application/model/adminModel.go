package model

import (
	"adminframe/framework/config"
	"adminframe/utils"
	"fmt"
	"time"
)

type AdminModel struct {
	ID int `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"-"`
	Salt string `gorm:"column:salt" json:"-"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Avatar string `gorm:"column:avatar" json:"avatar"`
	LastIP string `gorm:"column:last_ip" json:"last_ip"`
	IsSup int8 `gorm:"column:is_sup" json:"is_sup"`
	RoleID int `gorm:"column:role_id" json:"role_id"`
	LoginFailure int8 `gorm:"column:loginfailure" json:"login_failure"`
	LoginTime int `gorm:"column:logintime" json:"login_time"`
	CreateTime int `gorm:"column:create_time" json:"create_time"`
	UpdateTime int `gorm:"column:update_time" json:"update_time"`
}
// #### 设置表名
func (a AdminModel)TableName()string{
	return config.MysqlSetting.Prefix + "admin"
}

func NewAdminModel()*AdminModel{
	return new(AdminModel)
}
func (a *AdminModel)LoginAuth(pass string)error{
	_pass := utils.MD5Encry(utils.MD5Encry(pass) + a.Salt)
	if a.Password != _pass {
		if a.LoginFailure >= int8(config.AppSetting.LoginFailureTime) {
			a.LoginFailure = 1
		}else {
			a.LoginFailure += 1
		}
		a.UpdateTime = int(time.Now().Unix())
		var residue_times int8 = 0
		if a.LoginFailure < int8(config.AppSetting.LoginFailureTime) {
			residue_times = int8(config.AppSetting.LoginFailureTime) - a.LoginFailure
		}
		DB.Save(a)
		return fmt.Errorf("用户密码错误(剩余次数：%d)",residue_times)
	}
	a.LoginFailure = 0
	a.UpdateTime = int(time.Now().Unix())
	a.LoginTime = int(time.Now().Unix())
	return DB.Save(a).Error
}

func FindAdminByCondition(condition map[string]interface{}) (*AdminModel,error){
	whereSql,values,err := buildQuery(condition)
	if err != nil {
		return nil,err
	}
	model := new(AdminModel)
	err = DB.Where(whereSql,values...).Find(model).Error
	if err != nil {
		return nil,err
	}
	return model,nil
}