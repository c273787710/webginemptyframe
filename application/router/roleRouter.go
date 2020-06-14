package router

import (
	"github.com/gin-gonic/gin"
	"adminframe/middleware"
	"adminframe/application/controller/admin"
)

func roleRouter(e *gin.Engine){
	g := e.Group("/v1/admin/role")
	g.Use(middleware.JWTMiddleware(),middleware.AdminAuthMiddleware())
	{
		g.GET("/",admin.ListRole)
		g.POST("/",admin.AddRole)
		g.PUT("/",admin.UpdateRole)
		g.DELETE("/",admin.DelRole)
	}
}
