package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//日志中间件
func logMiddleWare(logger *zap.Logger)gin.HandlerFunc{
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		logger.Info(path,
			zap.String("path",path),
			zap.String("query",query),
			zap.String("method",c.Request.Method),
			zap.String("ip",c.ClientIP()),
			zap.Int("status",c.Writer.Status()),
			zap.String("errors",c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost",cost))
	}
}