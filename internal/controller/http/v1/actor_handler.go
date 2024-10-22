package v1

import (
	"net/http"
	"strconv"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/gin-gonic/gin"
)

type ActorHandler struct {
	actorService service.ActorService
}

func NewActorHandler(actorService service.ActorService) *ActorHandler {
	return &ActorHandler{actorService: actorService}
}

// @Summary Get all actors
// @Description Get all actors
// @Tags Actor
// @Produce  json
// @Router /actors [get]
// @Success 200 {object} model.HttpResponse[[]entity.Actor]
func (handler *ActorHandler) GetAll(c *gin.Context) {
	actors := handler.actorService.GetAllActor(c.Request.Context())
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[[]entity.Actor](&actors))
}

// @Summary Get an actor
// @Description Get an actor
// @Tags Actor
// @Produce  json
// @Param id path int true "actorId" example(1)
// @Router /actors/{id} [get]
// @Success 200 {object} model.HttpResponse[entity.Actor]
// @Failure 400 {object} model.HttpResponse[any]
// @Failure 500 {object} model.HttpResponse[any]
func (handler *ActorHandler) Get(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
	}

	actor, err := handler.actorService.GetActorById(c.Request.Context(), parsedId)
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
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.Actor](actor))
}

// @Summary Create an actor
// @Description Create an actor
// @Tags Actor
// @Accept json
// @Param params body model.ActorRequest true "Actor payload"
// @Produce  json
// @Router /actors [post]
// @Success 201 {object} model.HttpResponse[entity.Actor]
// @Failure 400 {object} model.HttpResponse[any]
// @Failure 500 {object} model.HttpResponse[any]
func (handler *ActorHandler) Create(c *gin.Context) {
	var actorRequest model.ActorRequest

	if err := c.ShouldBindJSON(&actorRequest); err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.HttpResponse[any]{Success: false, Data: nil})
		return
	}

	actor, err := handler.actorService.CreateActor(c.Request.Context(), actorRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[entity.Actor](actor))
}

// @Summary Update an actor
// @Description Update an actor
// @Tags Actor
// @Accept json
// @Param id path int true "actorId" example(1)
// @Param request body model.ActorRequest true "Actor payload"
// @Produce  json
// @Router /actors/{id} [put]
// @Success 200 {object} model.HttpResponse[entity.Actor]
// @Failure 400 {object} model.HttpResponse[any]
// @Failure 500 {object} model.HttpResponse[any]
func (handler *ActorHandler) Update(c *gin.Context) {
	//check param id
	id, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "", Field: "ID", Code: httpcommon.ErrorResponseCode.MissingIdParameter,
		}))
		return
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	var actorRequest model.ActorRequest
	//binding request
	if err = c.ShouldBindJSON(&actorRequest); err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.HttpResponse[any]{Success: false, Data: nil})
		return
	}
	//update
	updatedActor, err := handler.actorService.UpdateActor(c.Request.Context(), actorRequest, parsedId)
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
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.Actor](updatedActor))
}

// @Summary Delete an actor
// @Description Delete an actor with the given ID
// @Tags Actor
// @Produce json
// @Param id path int true "actorId" example(1)
// @Router /actors/{id} [delete]
// @Success 200 "Actor deleted successfully"
// @Failure 400 {object} model.HttpResponse[any]
// @Failure 500 {object} model.HttpResponse[any]
func (handler *ActorHandler) Delete(c *gin.Context) {
	//check param id
	id := c.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	err = handler.actorService.DeleteActor(c.Request.Context(), parsedId)
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
