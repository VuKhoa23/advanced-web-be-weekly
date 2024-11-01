package v1

import (
	"net/http"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/validation"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) Register(c *gin.Context) {
	var userRequest model.UserRequest

	if err := validation.BindJsonAndValidate(c, &userRequest); err != nil {
		return
	}

	user, err := handler.userService.Register(c.Request.Context(), userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[entity.User](user))
}

func (handler *UserHandler) Login(c *gin.Context) {
	var userRequest model.UserRequest

	if err := validation.BindJsonAndValidate(c, &userRequest); err != nil {
		return
	}

	token, err := handler.userService.Login(c.Request.Context(), userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}

	maxAge := 24 * 60 * 60 // 24h
	c.SetCookie("cookie", token, maxAge, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[string](&token))
}