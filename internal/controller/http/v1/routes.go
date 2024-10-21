package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
)

func LoggingRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		err := c.Request.Body.Close()
		if err != nil {
			c.Next()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var bodyJSON map[string]interface{}
		err = json.Unmarshal(bodyBytes, &bodyJSON)
		if err != nil {
			c.Next()
			return
		}

		logrus.WithFields(logrus.Fields{
			"method": c.Request.Method + "-request",
			"path":   c.Request.URL.Path,
			"query":  c.Request.URL.RawQuery,
			"body":   bodyJSON,
		}).Info()
		c.Next()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggingResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		statusCode := c.Writer.Status()
		if statusCode >= 0 {
			if statusCode < 400 {
				logrus.WithFields(logrus.Fields{
					"method": c.Request.Method + "-response",
					"path":   c.Request.URL.Path,
					"query":  c.Request.URL.RawQuery,
					"body":   blw.body.String(),
				}).Info()
			} else {
				logrus.WithFields(logrus.Fields{
					"method": c.Request.Method + "-response",
					"path":   c.Request.URL.Path,
					"query":  c.Request.URL.RawQuery,
					"body":   blw.body.String(),
				}).Error()
			}

		}
	}
}

func MapRoutes(router *gin.Engine, actorHandler *ActorHandler, filmHandler *FilmHandler) {
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    1, // megabytes
		MaxBackups: 1,
		MaxAge:     1, //days
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})

	router.Use(LoggingRequestMiddleware())
	router.Use(LoggingResponseMiddleware())
	v1 := router.Group("/api/v1")
	{
		actors := v1.Group("/actors")
		{
			actors.POST("/", actorHandler.Create)
			actors.GET("/", actorHandler.GetAll)
			actors.GET("/:id", actorHandler.Get)
			actors.PUT("/:id", actorHandler.Update)
			actors.DELETE("/:id", actorHandler.Delete)
		}

		films := v1.Group("/films")
		{
			films.GET("/:id", filmHandler.Get)
			films.DELETE("/:id", filmHandler.Delete)
			films.POST("/", filmHandler.Create)
			films.PUT("/:id", filmHandler.Update)
			films.GET("/", filmHandler.GetAll)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
