package grpc

import (
	"context"
	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/grpc/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	v1.UnimplementedFilmServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetAllFilms(ctx context.Context, in *v1.Empty) (*v1.NewFilmResponse, error) {
	return &v1.NewFilmResponse{FilmResponse: "hello demo"}, nil
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		return
	}

	server := grpc.NewServer()
	v1.RegisterFilmServer(server, s)
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("grpc server failed to serve: %v", err)
	}
}
