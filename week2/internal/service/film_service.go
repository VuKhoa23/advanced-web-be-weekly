package service

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type FilmService interface {
	GetFilmById(ctx context.Context, id int64) (*entity.Film, error)
	DeleteFilm(ctx context.Context, id int64) error
	GetAllFilms(ctx context.Context) []entity.Film
}
