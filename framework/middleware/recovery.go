package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

//异常中间件


func recoveryMiddleware(logger *zap.Logger)gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if err := recover();err != nil {
				var brokenPipe bool
				if se,ok := err.(*net.OpError);ok{
					if strings.Contains(strings.ToLower(se.Error()),"broken pipe") ||
						strings.Contains(strings.ToLower(se.Error()),"connection reset by peer"){
						brokenPipe = true
					}
				}
				httpRequest,_ := httputil.DumpRequest(c.Request,false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,zap.Any("error",err),zap.String("request",string(httpRequest)))
					c.Error(err.(error))
					c.Abort()
					return
				}
				logger.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}