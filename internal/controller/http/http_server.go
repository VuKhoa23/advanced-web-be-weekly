package http

import (
	"fmt"
	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	actorHandler *v1.ActorHandler
	filmHandler  *v1.FilmHandler
	authHandler  *v1.AuthHandler
}

func NewServer(actorHandler *v1.ActorHandler, filmHandler *v1.FilmHandler, authHandler *v1.AuthHandler) *Server {
	return &Server{actorHandler: actorHandler, filmHandler: filmHandler, authHandler: authHandler}
}

func (s *Server) Run() {
	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, s.actorHandler, s.filmHandler, s.authHandler)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
