package mw

import (	
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		ts := time.Now()
		latency := ts.Sub(start)

		logger.Info(
			"Request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Int("bodySize", c.Writer.Size()),
			zap.String("ip", c.ClientIP()),
			zap.String("userAgent", c.Request.UserAgent()),
			zap.Time("mvts", ts),
			zap.String("pmvts", ts.Format("2006-01-02T15:04:05-0700")),
			zap.Duration("latency", latency),
			zap.String("pLatency", latency.String()),
		)
	}
}
