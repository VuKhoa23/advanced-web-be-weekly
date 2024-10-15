package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/gin-gonic/gin"
)

type FilmHandler struct {
	filmService service.FilmService
}

func NewFilmHandler(filmService service.FilmService) *FilmHandler {
	return &FilmHandler{filmService: filmService}
}

// @Summary Get a film
// @Description Get a film with the given ID
// @Tags Film
// @Produce json
// @Param id path int true "filmId"
// @Router /films/{id} [get]
// @Success 200 {object} entity.Film
func (handler *FilmHandler) Get(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err.Error(), Data: nil})
	}

	film, err := handler.filmService.GetFilmById(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusOK, model.HttpResponse[any]{Message: err.Error(), Data: nil})
			return
		}
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[entity.Film]{Message: "Success", Data: film})
}

// @Summary Delete a film
// @Description Delete a film with the given ID
// @Tags Film
// @Produce json
// @Param id path int true "filmId"
// @Router /films/{id} [delete]
// @Success 204 "Film deleted successfully"
func (handler *FilmHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: "Invalid film ID", Data: nil})
		return
	}

	err = handler.filmService.DeleteFilm(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: "Film not found", Data: nil})
			return
		}
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: "Internal server error", Data: nil})
		return
	}

	// No content response on successful deletion (204)
	c.Status(http.StatusNoContent)
}

// @Summary Create a film
// @Description Create a film
// @Tags Film
// @Accept json
// @Param  params body model.FilmRequest true "Film payload"
// @Produce  json
// @Router /films [post]
// @Success 201 {object} entity.Film
func(handler *FilmHandler) Create(c *gin.Context){
	var filmRequest model.FilmRequest

	if err := c.ShouldBindJSON(&filmRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	} 

	film := &entity.Film{
		Title:            filmRequest.Title,
		Description:      filmRequest.Description,
		ReleaseYear:      filmRequest.ReleaseYear,
		LanguageID:       filmRequest.LanguageID,
		OriginalLanguageID: filmRequest.OriginalLanguageID,
		RentalDuration:   filmRequest.RentalDuration,
		RentalRate:       filmRequest.RentalRate,
		Length:           filmRequest.Length,
		ReplacementCost:  filmRequest.ReplacementCost,
		Rating:           filmRequest.Rating,
		SpecialFeatures:  filmRequest.SpecialFeatures,
	}
	err := handler.filmService.CreateFilm(c.Request.Context(), film)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[entity.Film]{Message: "Success", Data: film})
}

// @Summary Update a film
// @Description Update a film
// @Tags Film
// @Accept json
// @Param  params body model.FilmRequest true "Film payload"
// @Produce  json
// @Router /films [put]
// @Success 200 {object} entity.Film
func(handler *FilmHandler) Update(c *gin.Context){
	filmId, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required").Error()})
		return
	}

	parseId, err := strconv.Atoi(filmId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var filmRequest model.FilmRequest
	if err := c.ShouldBindJSON(&filmRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	} 

	film := &entity.Film{
		ID:            	int64(parseId),
		Title:            filmRequest.Title,
		Description:      filmRequest.Description,
		ReleaseYear:      filmRequest.ReleaseYear,
		LanguageID:       filmRequest.LanguageID,
		OriginalLanguageID: filmRequest.OriginalLanguageID,
		RentalDuration:   filmRequest.RentalDuration,
		RentalRate:       filmRequest.RentalRate,
		Length:           filmRequest.Length,
		ReplacementCost:  filmRequest.ReplacementCost,
		Rating:           filmRequest.Rating,
		SpecialFeatures:  filmRequest.SpecialFeatures,
	}
	err = handler.filmService.UpdateFilm(c.Request.Context(), film)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[entity.Film]{Message: "Success", Data: film})
}

