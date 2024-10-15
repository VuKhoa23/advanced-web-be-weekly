package serviceimplement

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type FilmService struct {
	filmRepository repository.FilmRepository
}

func NewFilmService(filmRepository repository.FilmRepository) service.FilmService {
	return &FilmService{filmRepository: filmRepository}
}

func (service *FilmService) GetFilmById(ctx context.Context, id int64) (*entity.Film, error) {
	return service.filmRepository.GetFilmByID(ctx, id)
}

func (service *FilmService) DeleteFilm(ctx context.Context, id int64) error {
	result := service.filmRepository.DeleteFilm(ctx, id)
	if result != nil {
		return result
	}
	return nil
}

func (service *FilmService) GetAllFilms(ctx context.Context) []entity.Film {
	return service.filmRepository.GetAllFilms(ctx)
}
