package model
//数据库连接配置
import (
	"adminframe/framework/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var DB *gorm.DB
func InitModel(){
	var err error
	DB,err = gorm.Open("mysql",dns())
	if err != nil {
		panic(fmt.Sprintf("Connect Mysql Error : %s",err))
	}
	DB.DB().SetMaxOpenConns(config.MysqlSetting.MaxOpenConn)
	DB.DB().SetMaxIdleConns(config.MysqlSetting.MaxIdleConn)
}

func dns()string{
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.MysqlSetting.Username,
		config.MysqlSetting.Password,
		config.MysqlSetting.Host,
		config.MysqlSetting.Port,
		config.MysqlSetting.Databases,
		config.MysqlSetting.Charset)
	return dns
}