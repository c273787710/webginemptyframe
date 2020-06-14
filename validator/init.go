package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitSelfValidator(){
	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("rulepidvalid",rulePidValidator)
		v.RegisterValidation("ruleidsisvalid",ruleIDIsValidator)
		v.RegisterValidation("roleidisvalid",roleidisvalid)
		v.RegisterValidation("rulemethodvalidate",requestmethodvalid)
	}
}
