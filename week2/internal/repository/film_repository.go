package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type FilmRepository interface {
	GetFilmByID(c context.Context, id int64) (*entity.Film, error)
	DeleteFilm(c context.Context, id int64) error
	GetAllFilms(c context.Context) []entity.Film
}
