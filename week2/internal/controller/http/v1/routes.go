package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/logdyhq/logdy-core/logdy"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func Middleware(logger logdy.Logdy) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyData interface{}
		if err := c.ShouldBindBodyWith(&bodyData, binding.JSON); err != nil {
			print(err.Error())
			if err.Error() == "EOF" {
				bodyData = ""
			}
		}
		logger.Log(logdy.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"query":  c.Request.URL.RawQuery,
			"body":   bodyData,
			"time":   time.Now(),
		})
		c.Next()
	}
}

func MapRoutes(router *gin.Engine, logger logdy.Logdy, actorHandler *ActorHandler, filmHandler *FilmHandler) {
	router.Use(Middleware(logger))
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
