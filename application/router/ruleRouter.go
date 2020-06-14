package router

import (
	"github.com/gin-gonic/gin"
	"adminframe/application/controller/admin"
	"adminframe/middleware"
)

func ruleRouter(e *gin.Engine){
	g := e.Group("/v1/admin/rule")
	g.Use(middleware.JWTMiddleware(),middleware.AdminAuthMiddleware())
	{
		g.POST("/",admin.AddRule)
		g.PUT("/",admin.UpdateRule)
		g.GET("/",admin.ListRule)
		g.DELETE("/",admin.DelRule)
	}
}
