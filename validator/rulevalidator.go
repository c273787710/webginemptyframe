package validator

import (
	"github.com/go-playground/validator/v10"
	"adminframe/application/model"
	"strings"
)

func rulePidValidator(field validator.FieldLevel)bool{
	if data,ok := field.Field().Interface().(int);ok{
		if data == 0 {
			return true
		}
		query := map[string]interface{}{
			"id": data,
		}
		model,err := model.FindRuleByCondition(query)
		if err != nil || model == nil {
			return false
		}
		return true
	}
	return false
}

func ruleIDIsValidator(field validator.FieldLevel)bool{
	if data,ok := field.Field().Interface().([]int);ok{
		for _,v := range data {
			exit,_ := model.FindRuleByCondition(map[string]interface{}{"id":v})
			if exit == nil {
				return false
			}
		}
		return true
	}
	return false
}

var methods = []string{"*","get","post","delete","put"}
func requestmethodvalid(field validator.FieldLevel)bool{
	if data,ok := field.Field().Interface().(string);ok{
		if data == "" {
			return true
		}
		for _,v := range methods {
			if v == strings.ToLower(data) {
				return true
			}
		}
		return false
	}
	return false
}