package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/logdyhq/logdy-core/logdy"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
)

type Server struct {
	actorHandler *v1.ActorHandler
	filmHandler  *v1.FilmHandler
}

func NewServer(actorHandler *v1.ActorHandler, filmHandler *v1.FilmHandler) *Server {
	return &Server{actorHandler: actorHandler, filmHandler: filmHandler}
}

func (s *Server) Run() {
	// setup logrus
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	logrus.SetLevel(logrus.InfoLevel)

	// Log 10,000 entries
	for i := 0; i < 10000; i++ {
		logrus.Infof("This is log entry number %d", i)
	}

	// setup logdy
	logger := logdy.InitializeLogdy(logdy.Config{
		ServerIp:   "127.0.0.1",
		ServerPort: "8081",
	}, nil)

	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, logger, s.actorHandler, s.filmHandler)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
