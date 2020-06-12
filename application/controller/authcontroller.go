package controller

import (
	"adminframe/utils"
	"github.com/gin-gonic/gin"
)

func LoginAuth(c *gin.Context){
	object := utils.Object{C:c}
	object.Response(200,nil,"hello world")
	return
}
