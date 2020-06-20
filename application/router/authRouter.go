package router

import (
	"adminframe/application/controller/admin"
	"adminframe/middleware"
	"github.com/gin-gonic/gin"
)

//登录授权相关路由

func authRouter(e *gin.Engine){
	//#### 登录路由
	e.POST("/v1/admin/auth", admin.LoginAuth)
	g := e.Group("/v1/admin")
	g.Use(middleware.JWTMiddleware())
	{
		g.GET("/getInfo",admin.AdminMenu)
	}
}
