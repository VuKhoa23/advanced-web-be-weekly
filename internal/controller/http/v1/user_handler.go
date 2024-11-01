package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
func (handler *UserHandler) Register(c *gin.Context) {
	var userRequest model.UserRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	userName, err := handler.userService.CreateUser(c.Request.Context(), userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[entity.User](&entity.User{UserName: userName}))
}

func (handler *UserHandler) Login(c *gin.Context) {
	var userRequest model.UserRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	user, err := handler.userService.CheckPassword(c.Request.Context(), userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InvalidUserInfo,
		}))
		return
	}
	//create token
	tokenString, err := authentication.GenerateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
	}
	//setup cookie
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    tokenString,
		HttpOnly: true,
		Expires:  time.Now().Add(3 * time.Hour),
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[entity.User](&entity.User{UserName: user.UserName}))
}
