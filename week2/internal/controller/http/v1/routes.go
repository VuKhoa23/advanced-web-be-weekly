package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/logdyhq/logdy-core/logdy"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, logdyLogger logdy.Logdy, actorHandler *ActorHandler, filmHandler *FilmHandler) {
	router.Use(middleware.Logging(logdyLogger))
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
