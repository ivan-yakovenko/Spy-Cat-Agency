package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startingTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(startingTime)
		logrus.WithFields(logrus.Fields{
			"method":   method,
			"path":     path,
			"status":   status,
			"duration": duration,
		}).Info("Completed request")

	}
}
