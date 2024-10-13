package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
)

type Server struct {
	actorHandler *v1.ActorHandler
}

func NewServer(actorHandler *v1.ActorHandler) *Server {
	return &Server{actorHandler: actorHandler}
}

func (s *Server) Run() {
	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, s.actorHandler)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
