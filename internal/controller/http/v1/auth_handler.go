package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/authentication"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/constants"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (handler *AuthHandler) Register(c *gin.Context) {
	var registerRequest model.RegisterRequest

	if err := c.ShouldBind(&registerRequest); err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	username, err := handler.userService.Register(c.Request.Context(), registerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[entity.User](&entity.User{Username: username}))
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var loginRequest model.LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	user, err := handler.userService.Login(c.Request.Context(), loginRequest)
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
		Expires:  time.Now().Add(constants.COOKIE_DURATION),
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[entity.User](&entity.User{Username: user.Username}))
}
