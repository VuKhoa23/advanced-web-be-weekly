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

func (service *FilmService) CreateFilm(ctx context.Context, film *entity.Film) error {
	result := service.filmRepository.CreateFilm(ctx, film)
	if result != nil {
		return result
	}
	return nil
}

func (service *FilmService) UpdateFilm(ctx context.Context, film *entity.Film) error {
	result := service.filmRepository.UpdateFilm(ctx, film)
	if result != nil {
		return result
	}
	return nil
}

