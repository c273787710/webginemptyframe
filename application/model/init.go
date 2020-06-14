package model
//数据库连接配置
import (
	"adminframe/framework/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
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
	DB.DB().SetConnMaxLifetime(time.Hour)
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

func buildQuery(where map[string]interface{})(whereSql string,values []interface{},err error){
	for key,val := range where{
		conditionKey := strings.Split(key," ")
		if len(conditionKey) > 2 {
			return "",nil,fmt.Errorf("查询条件出错")
		}
		if whereSql != "" {
			whereSql += " AND "
		}
		switch len(conditionKey) {
		case 1:
			whereSql +=fmt.Sprint(conditionKey[0]," = ?")
			values = append(values,val)
			break
		case 2:
			field := conditionKey[0]
			values = append(values,val)
			switch conditionKey[1] {
			case "=":
				whereSql += fmt.Sprint(field," = ?")
				break
			case ">":
				whereSql += fmt.Sprint(field," > ?")
				break
			case ">=":
				whereSql += fmt.Sprint(field," >= ?")
				break
			case "<":
				whereSql += fmt.Sprint(field," < ?")
				break
			case "<=":
				whereSql += fmt.Sprint(field," <= ?")
				break
			case "<>":
				whereSql += fmt.Sprint(field," != ?")
				break
			case "!=":
				whereSql += fmt.Sprint(field," != ?")
				break
			case "in":
				whereSql += fmt.Sprint(field," in (?)")
				break
			case "like":
				whereSql += fmt.Sprint(field," like ?")
				break
			}
		}
	}
	return
}