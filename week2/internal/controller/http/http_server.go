package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening log file: %v", err)
	}
	defer logFile.Close()

	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.InfoLevel)

	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, s.actorHandler, s.filmHandler)
	err = httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
