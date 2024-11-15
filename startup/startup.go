package startup

import (
	"github.com/VuKhoa23/advanced-web-be/internal"
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/validation"
	"github.com/gammazero/workerpool"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	return internal.InitializeContainer(db)
}

func startServers(container *controller.ApiContainer) {
	wp := workerpool.New(2)

	wp.Submit(container.HttpServer.Run)
	wp.Submit(container.GrpcServer.Run)

	wp.StopWait()
}

func Execute() {
	validation.GetValidations()
	container := registerDependencies()
	startServers(container)
}
