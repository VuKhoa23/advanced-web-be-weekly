package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/validation"
	"github.com/gin-gonic/gin"
	v1 "github.com/hoadang0305/grpc-server-b/public/controller/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
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
// @Param id path int true "filmId" example(1)
// @Router /films/{id} [get]
// @Success 200 {object} httpcommon.HttpResponse[entity.Film]
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *FilmHandler) Get(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
	}

	film, err := handler.filmService.GetFilmById(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.GormRecordNotFound {
			c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
				Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.RecordNotFound,
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.Film](film))
}

// @Summary Delete a film
// @Description Delete a film with the given ID
// @Tags Film
// @Produce json
// @Param id path int true "filmId" example(1)
// @Router /films/{id} [delete]
// @Success 200 "Film deleted successfully"
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *FilmHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	err = handler.filmService.DeleteFilm(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.GormRecordNotFound {
			c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
				Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.RecordNotFound,
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}

	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[any](nil))
}

// @Summary Create a film
// @Description Create a film
// @Tags Film
// @Accept json
// @Param request body model.FilmRequest true "Film payload"
// @Produce  json
// @Router /films [post]
// @Success 200 {object} httpcommon.HttpResponse[entity.Film]
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *FilmHandler) Create(c *gin.Context) {
	var filmRequest model.FilmRequest

	if err := validation.BindJsonAndValidate(c, &filmRequest); err != nil {
		return
	}

	film, err := handler.filmService.CreateFilm(c.Request.Context(), filmRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.Film](film))
}

// @Summary Update a film
// @Description Update a film
// @Tags Film
// @Accept json
// @Param id path int true "filmId" example(1)
// @Param request body model.FilmRequest true "Film payload"
// @Produce  json
// @Router /films [put]
// @Success 200 {object} httpcommon.HttpResponse[entity.Film]
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *FilmHandler) Update(c *gin.Context) {
	filmId, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "", Field: "ID", Code: httpcommon.ErrorResponseCode.MissingIdParameter,
		}))
		return
	}

	parsedId, err := strconv.ParseInt(filmId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	var filmRequest model.FilmRequest
	if err = validation.BindJsonAndValidate(c, &filmRequest); err != nil {
		return
	}

	film, err := handler.filmService.UpdateFilm(c.Request.Context(), filmRequest, parsedId)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.GormRecordNotFound {
			c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
				Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.RecordNotFound,
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.Film](film))
}

// @Summary Get all films
// @Description Get all films
// @Tags Film
// @Produce  json
// @Router /films [get]
// @Success 200 {object} httpcommon.HttpResponse[[]entity.Film]
func (handler *FilmHandler) GetAll(c *gin.Context) {
	address := "localhost:3001"
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}))
	}

	defer conn.Close()
	connClient := v1.NewFilmHandlerClient(conn)
	res, err := connClient.GetAllFilms(c, &v1.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}))
	}

	films := []v1.Film{}
	for _, film := range res.GetListfilms() {
		films = append(films, *film)
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[[]v1.Film](&films))
}
