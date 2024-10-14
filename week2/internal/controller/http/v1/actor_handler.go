package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
// @Success 200 {object} []entity.Actor
func (handler *ActorHandler) GetAll(c *gin.Context) {
	actors := handler.actorService.GetAllActor(c.Request.Context())
	c.JSON(http.StatusOK, model.HttpResponse[[]entity.Actor]{Message: "Success", Data: &actors})
}

// @Summary Get an actor
// @Description Get an actor
// @Tags Actor
// @Produce  json
// @Param id path int true "actorId"
// @Router /actors/{id} [get]
// @Success 200 {object} entity.Actor
func (handler *ActorHandler) Get(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err.Error(), Data: nil})
	}

	actor, err := handler.actorService.GetActorById(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusOK, model.HttpResponse[any]{Message: err.Error(), Data: nil})
			return
		}
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[entity.Actor]{Message: "Success", Data: actor})
}

// @Summary Create an actor
// @Description Create an actor
// @Tags Actor
// @Accept json
// @Param  params body model.ActorRequest true "Actor payload"
// @Produce  json
// @Router /actors [post]
// @Success 200 {object} entity.Actor
func (handler *ActorHandler) Create(c *gin.Context) {
	var actorRequest model.ActorRequest

	if err := c.ShouldBindJSON(&actorRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	}

	actor := &entity.Actor{FirstName: actorRequest.FirstName, LastName: actorRequest.LastName}
	err := handler.actorService.CreateActor(c.Request.Context(), actor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[entity.Actor]{Message: "Success", Data: actor})
}

// @Summary Update an actor
// @Description Update an actor
// @Tags Actor
// @Accept json
// @Param id path int true "actorId"
// @RequestBody model.ActorRequest true "Actor payload"
// @Produce  json
// @Router /actors/{id} [put]
// @Success 200 {object} entity.Actor
func (handler *ActorHandler) Update(c *gin.Context) {
	//check param id
	id := c.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err.Error(), Data: nil})
	}

	var actorRequest model.ActorRequest
	//binding request
	if err1 := c.ShouldBindJSON(&actorRequest); err1 != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err1.Error(), Data: nil})
		return
	}
	//update
	updatedActor := &entity.Actor{ID: parsedId, FirstName: actorRequest.FirstName, LastName: actorRequest.LastName}
	err2 := handler.actorService.UpdateActor(c.Request.Context(), updatedActor)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: err2.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[entity.Actor]{Message: "Success", Data: updatedActor})
}

// @Summary Delete an actor
// @Description Delete an actor with the given ID
// @Tags Actor
// @Produce json
// @Param id path int true "actorId"
// @Router /actors/{id} [delete]
// @Success 204 "Actor deleted successfully"
func (handler *ActorHandler) Delete(c *gin.Context) {
	//check param id
	id := c.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: err.Error(), Data: nil})
	}

	err = handler.actorService.DeleteActor(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, model.HttpResponse[any]{Message: "Actor not found", Data: nil})
			return
		}
		c.JSON(http.StatusInternalServerError, model.HttpResponse[any]{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[any]{Message: "Actor deleted successfully", Data: nil})
}
