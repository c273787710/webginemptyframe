package router

import (
	"adminframe/application/controller"
	"github.com/gin-gonic/gin"
)

//登录授权相关路由

func authRouter(e *gin.Engine){
	e.GET("/v1/admin/auth", controller.LoginAuth)
}
