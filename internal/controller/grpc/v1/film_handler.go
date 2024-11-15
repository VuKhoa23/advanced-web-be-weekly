package v1

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type FilmHandler struct {
	filmService service.FilmService
}

func NewFilmHandler(filmService service.FilmService) *FilmHandler {
	return &FilmHandler{filmService: filmService}
}

func (handler *FilmHandler) GetAll(c context.Context) *[]entity.Film {
	films := handler.filmService.GetAllFilms(c)
	return &films
}
