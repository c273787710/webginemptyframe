package validator

import (
	"github.com/go-playground/validator/v10"
	"adminframe/application/model"
)

func roleidisvalid(field validator.FieldLevel)bool{
	if data,ok := field.Field().Interface().(int);ok {
		if data == 0 {
			return true
		}
		exit,_ := model.FindRoleByCondition(map[string]interface{}{"id":data})
		if exit == nil {
			return false
		}
		return true
	}
	return false
}
