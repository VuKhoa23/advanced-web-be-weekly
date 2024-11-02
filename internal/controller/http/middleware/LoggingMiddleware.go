package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

func LoggingRequestMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.Next()
		}
	}()

	bodyBytes, _ := io.ReadAll(c.Request.Body)
	// close request body to reuse underlying TCP socket
	_ = c.Request.Body.Close()
	// re populate the Body
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	var bodyJSON map[string]interface{}
	err := json.Unmarshal(bodyBytes, &bodyJSON)
	if err != nil {
		bodyJSON = nil
	}

	logrus.WithFields(logrus.Fields{
		"method": c.Request.Method + "-request",
		"path":   c.Request.URL.Path,
		"body":   bodyJSON,
	}).Info()
	c.Next()
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggingResponseMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.Next()
		}
	}()

	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	if statusCode >= 0 {
		if statusCode < 400 {
			logrus.WithFields(logrus.Fields{
				"method": c.Request.Method + "-response",
				"path":   c.Request.URL.Path,
				"body":   blw.body.String(),
			}).Info()
		} else {
			logrus.WithFields(logrus.Fields{
				"method": c.Request.Method + "-response",
				"path":   c.Request.URL.Path,
				"body":   blw.body.String(),
			}).Error()
		}

	}
}
