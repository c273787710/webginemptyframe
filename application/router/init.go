package router

import (
	"github.com/gin-gonic/gin"
)

var routers []func(e *gin.Engine)

func register(r func(e *gin.Engine)){
	routers = append(routers,r)
}

func InitRouter(r *gin.Engine){
	register(authRouter)
	for _,router := range routers{
		router(r)
	}
}
