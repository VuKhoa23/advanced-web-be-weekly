package startup

import (
	"github.com/VuKhoa23/advanced-web-be/internal"
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	// migration
	err := db.AutoMigrate(&entity.Actor{})
	if err != nil {
		panic(err)
	}

	return internal.InitializeContainer(db)
}

func Execute() {
	container := registerDependencies()
	container.HttpServer.Run()
}
