package v1

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func MapRoutes(router *gin.Engine, actorHandler *ActorHandler, filmHandler *FilmHandler) {
	currentTime := time.Now()
	formattedDate := currentTime.Format("02-01-2006")

	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename: "logs/" + formattedDate + ".log",
		MaxSize:  10, // megabytes
		MaxAge:   7,  // days
	})
	logrus.SetOutput(multiWriter)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	router.Use(gin.Recovery())
	router.Use(LoggingRequestMiddleware)
	router.Use(LoggingResponseMiddleware)
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
			films.GET("/", filmHandler.GetAll)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
