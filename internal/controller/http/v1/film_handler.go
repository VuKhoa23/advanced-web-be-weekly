package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/VuKhoa23/advanced-web-be/internal/constants"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/authentication"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/validation"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/gin-gonic/gin"
)

type FilmHandler struct {
	filmService service.FilmService
	kafkaProducer service.KafkaService
}

func NewFilmHandler(filmService service.FilmService, kafkaProducer service.KafkaService) *FilmHandler {
	return &FilmHandler{
		filmService: filmService,
		kafkaProducer: kafkaProducer,
	}
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

	// Serialize the film request
	reqBody, err := json.Marshal(filmRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Code: httpcommon.ErrorResponseCode.InternalServerError, Field: "",
		}))
		return
	}

	// Send request to Kafka
	reply, err := handler.kafkaProducer.SendMessage(constants.REQUEST_TOPIC, string(reqBody)); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Code: httpcommon.ErrorResponseCode.InternalServerError, Field: "",
		}))
		return
	}

	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[string](&reply))
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
	//films := handler.filmService.GetAllFilms(c.Request.Context())

	//create client and setup
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3001/api/v1/films", nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}))
		return
	}

	// generate a token from API key to include in request
	requestTime := time.Now().Unix()
	token, err := authentication.GenerateTokenFromApiKey("http://localhost:3001/api/v1/films", requestTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}))
		return
	}
	req.Header.Set("Token", token)
	req.Header.Set("Request-Time", fmt.Sprintf("%d", requestTime))

	//call API and handler response
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}))
	}
	//mapping and response
	var res httpcommon.HttpResponse[[]entity.Film]
	if err := json.Unmarshal(body, &res); err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}))
	}
	c.JSON(http.StatusOK, res)
}
