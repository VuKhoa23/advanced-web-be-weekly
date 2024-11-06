// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http"
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/repository/implement"
	"github.com/VuKhoa23/advanced-web-be/internal/service/implement"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeContainer(db database.Db) *controller.ApiContainer {
	actorRepository := repositoryimplement.NewActorRepository(db)
	actorService := serviceimplement.NewActorService(actorRepository)
	actorHandler := v1.NewActorHandler(actorService)
	filmRepository := repositoryimplement.NewFilmRepository(db)
	filmService := serviceimplement.NewFilmService(filmRepository)
	filmHandler := v1.NewFilmHandler(filmService)
	userRepository := repositoryimplement.NewUserRepository(db)
	userService := serviceimplement.NewUserService(userRepository)
	refreshTokenRepository := repositoryimplement.NewRefreshTokenRepository(db)
	refreshTokenService := serviceimplement.NewRefreshTokenService(refreshTokenRepository)
	authHandler := v1.NewAuthHandler(userService, refreshTokenService)
	server := http.NewServer(actorHandler, filmHandler, authHandler)
	apiContainer := controller.NewApiContainer(server)
	return apiContainer
}

// wire.go:

var container = wire.NewSet(controller.NewApiContainer)

// may have grpc server in the future
var serverSet = wire.NewSet(http.NewServer)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(v1.NewActorHandler, v1.NewFilmHandler, v1.NewAuthHandler)

var serviceSet = wire.NewSet(serviceimplement.NewActorService, serviceimplement.NewFilmService, serviceimplement.NewUserService, serviceimplement.NewRefreshTokenService)

var repositorySet = wire.NewSet(repositoryimplement.NewActorRepository, repositoryimplement.NewFilmRepository, repositoryimplement.NewUserRepository, repositoryimplement.NewRefreshTokenRepository)
