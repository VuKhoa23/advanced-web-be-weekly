package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context){
		logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"query":  c.Request.URL.RawQuery,
			"time":   c.Request.Header.Get("X-Request-Time"),
		}).Info("Request details")
		c.Next()
	}
}