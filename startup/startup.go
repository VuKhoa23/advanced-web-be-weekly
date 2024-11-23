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
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	// Initialize FilmService and FilmHandler
    filmRepository := repositoryimplement.NewFilmRepository(db)
    filmService := serviceimplement.NewFilmService(filmRepository)
    filmHandler := v1.NewFilmHandler(filmService)

    // Create a Kafka message to handle film creation
    kafkaMessageHandler := func(ctx context.Context, message []byte) error {
        return filmHandler.Create(ctx, message)
    }

	// Initialize Kafka Consumer Service
	kafkaService := serviceimplement.NewKafkaService([]string{constants.BROKER}, kafkaMessageHandler)

	// Start consuming Kafka messages
	if kafkaService != nil {
		topics := []string{constants.REQUEST_TOPIC}
		go kafkaService.Start(context.Background(), topics)
		log.Println("Kafka consumer service started")
	} else {
		log.Println("Failed to initialize Kafka consumer service")
	}


	return internal.InitializeContainer(db)
}

func Execute() {
	validation.GetValidations()

	container := registerDependencies()
	container.HttpServer.Run()
}
