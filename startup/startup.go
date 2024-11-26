package startup

import (
	"context"
	"log"

	"github.com/VuKhoa23/advanced-web-be/internal"
	"github.com/VuKhoa23/advanced-web-be/internal/constants"
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	repositoryimplement "github.com/VuKhoa23/advanced-web-be/internal/repository/implement"
	serviceimplement "github.com/VuKhoa23/advanced-web-be/internal/service/implement"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/validation"
	"github.com/gammazero/workerpool"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	// Initialize FilmService and FilmHandler
    filmRepository := repositoryimplement.NewFilmRepository(db)
    filmService := serviceimplement.NewFilmService(filmRepository)
    filmHandler := v1.NewFilmHandler(filmService)

    // Create a Kafka message handler that processes film creation
    kafkaMessageHandler := func(ctx context.Context, message []byte) error {
        return filmHandler.Create(ctx, message)
    }

	// Initialize Kafka Consumer Service
	kafkaService := serviceimplement.NewKafkaService([]string{constants.BROKER}, kafkaMessageHandler) // Replace with your broker addresses

	// Start consuming Kafka messages
	if kafkaService != nil {
		topics := []string{constants.TOPIC}
		go kafkaService.Start(context.Background(), topics)
		log.Println("Kafka consumer service started")
	} else {
		log.Println("Failed to initialize Kafka consumer service")
	}

	return internal.InitializeContainer(db)
}

func startServers(container *controller.ApiContainer) {
	wp := workerpool.New(1)

	wp.Submit(container.HttpServer.Run)

	wp.StopWait()
}

func Execute() {
	validation.GetValidations()

	container := registerDependencies()
	startServers(container)
}
