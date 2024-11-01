package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
	"time"
)

func MapRoutes(router *gin.Engine, actorHandler *ActorHandler, filmHandler *FilmHandler, authHandler *AuthHandler) {
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
	router.Use(middleware.LoggingRequestMiddleware)
	router.Use(middleware.LoggingResponseMiddleware)
	v1 := router.Group("/api/v1")
	{
		actors := v1.Group("/actors")
		{
			actors.POST("/", actorHandler.Create)
			actors.GET("/", middleware.VerifyTokenMiddleware, actorHandler.GetAll)
			actors.GET("/:id", actorHandler.Get)
			actors.PUT("/:id", actorHandler.Update)
			actors.DELETE("/:id", actorHandler.Delete)
		}
		films := v1.Group("/films")
		{
			films.GET("/", middleware.VerifyTokenMiddleware, filmHandler.GetAll)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
