package middleware

import (
	"adminframe/framework/utils"
	"github.com/gin-gonic/gin"
)

func InitMiddleWare(e *gin.Engine){
	logger := utils.NewLogger()
	//跨域访问中间件
	e.Use(corsMiddleWare())
	//日志记录中间件
	e.Use(logMiddleWare(logger))
	//异常恢复中间件
	e.Use(recoveryMiddleware(logger))
}
