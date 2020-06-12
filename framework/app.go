package framework

import (
	"adminframe/application/router"
	"adminframe/framework/middleware"
	"github.com/gin-gonic/gin"
)

func NewApp()*gin.Engine{
	r := gin.New()
	//注册中间件
	middleware.InitMiddleWare(r)
	//注册路由
	router.InitRouter(r)
	return r
}
