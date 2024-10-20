package serviceimplement

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"

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

func (service *FilmService) CreateFilm(ctx context.Context, filmRequest model.FilmRequest) (*entity.Film, error) {
	film := &entity.Film{
		Title:              filmRequest.Title,
		Description:        filmRequest.Description,
		ReleaseYear:        filmRequest.ReleaseYear,
		LanguageID:         filmRequest.LanguageId,
		OriginalLanguageID: filmRequest.OriginalLanguageId,
		RentalDuration:     filmRequest.RentalDuration,
		RentalRate:         filmRequest.RentalRate,
		Length:             filmRequest.Length,
		ReplacementCost:    filmRequest.ReplacementCost,
		Rating:             filmRequest.Rating,
		SpecialFeatures:    filmRequest.SpecialFeatures,
	}

	err := service.filmRepository.CreateFilm(ctx, film)
	if err != nil {
		return nil, err
	}
	return film, nil
}

func (service *FilmService) UpdateFilm(ctx context.Context, filmRequest model.FilmRequest, filmId int64) (*entity.Film, error) {
	film := &entity.Film{
		Title:              filmRequest.Title,
		Description:        filmRequest.Description,
		ReleaseYear:        filmRequest.ReleaseYear,
		LanguageID:         filmRequest.LanguageId,
		OriginalLanguageID: filmRequest.OriginalLanguageId,
		RentalDuration:     filmRequest.RentalDuration,
		RentalRate:         filmRequest.RentalRate,
		Length:             filmRequest.Length,
		ReplacementCost:    filmRequest.ReplacementCost,
		Rating:             filmRequest.Rating,
		SpecialFeatures:    filmRequest.SpecialFeatures,
	}

	updatedFilm, err := service.filmRepository.UpdateFilm(ctx, film, filmId)
	if err != nil {
		return nil, err
	}
	return updatedFilm, nil
}
func (service *FilmService) GetAllFilms(ctx context.Context) []entity.Film {
	return service.filmRepository.GetAllFilms(ctx)
}
