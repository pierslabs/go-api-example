package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger, _ := zap.NewProduction()

		defer logger.Sync()
		logger.Info("TestMiddleware executed",
			zap.String("url", c.Request.URL.Path),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
		c.Next()
	}
}
