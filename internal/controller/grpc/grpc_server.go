package grpc

import (
	"context"
	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

type Server struct {
	v1.UnimplementedFilmHandlerServer
	filmHandler *v1.FilmHandler
}

func NewServer(filmHandler *v1.FilmHandler) *Server {
	return &Server{filmHandler: filmHandler}
}

func (s *Server) GetAllFilms(ctx context.Context, in *v1.Empty) (*v1.ListFilmResponse, error) {
	films := s.filmHandler.GetAll(ctx)
	filmList := &v1.ListFilmResponse{}
	for _, film := range *films {
		filmList.Listfilms = append(filmList.Listfilms, &v1.Film{
			Id:                 film.Id,
			Title:              film.Title,
			Description:        film.Description,
			ReleaseYear:        uint32(film.ReleaseYear),
			LanguageId:         film.LanguageID,
			OriginalLanguageId: film.OriginalLanguageID,
			RentalDuration:     film.RentalDuration,
			RentalRate:         film.RentalRate,
			Length:             film.Length,
			ReplacementCost:    film.ReplacementCost,
			Rating:             film.Rating,
			SpecialFeatures:    film.SpecialFeatures,
			LastUpdate:         timestamppb.New(film.LastUpdate),
		})
	}
	return filmList, nil
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		return
	}

	server := grpc.NewServer()
	v1.RegisterFilmHandlerServer(server, s)
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("grpc server failed to serve: %v", err)
	}
}
