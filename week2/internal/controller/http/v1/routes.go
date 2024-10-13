package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, actorHandler *ActorHandler) {
	v1 := router.Group("/api/v1")
	{
		actors := v1.Group("/actors")
		{
			actors.POST("/", actorHandler.Create)
			actors.GET("/", actorHandler.GetAll)
			actors.GET("/:id", actorHandler.Get)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
