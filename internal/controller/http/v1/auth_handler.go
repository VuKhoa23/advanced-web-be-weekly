package v1

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/authentication"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

type AuthHandler struct {
	userService         service.UserService
	refreshTokenService service.RefreshTokenService
}

func NewAuthHandler(userService service.UserService, refreshTokenService service.RefreshTokenService) *AuthHandler {
	return &AuthHandler{userService: userService, refreshTokenService: refreshTokenService}
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
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.User](&entity.User{Username: username}))
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	var refreshTokenRequest model.RefreshTokenRequest

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
	//create access token
	accessTokenExpTime := time.Now().Add(constants.JWT_DURATION)
	accessToken, err := authentication.GenerateToken(user, accessTokenExpTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
	}

	//create refresh token
	refreshTokenExpTime := time.Now().Add(constants.REFRESH_TOKEN_DURATION)
	fmt.Println("exp time ", refreshTokenExpTime)
	refreshToken, err := authentication.GenerateToken(user, refreshTokenExpTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
	}

	refreshTokenRequest = model.RefreshTokenRequest{
		Token:    refreshToken,
		Username: user.Username,
		ExpTime:  refreshTokenExpTime,
	}

	// check refresh token in database to see if it exist
	storedRefreshToken, err := handler.refreshTokenService.FindRefreshTokenByUsername(c.Request.Context(), user.Username) // find by username
	if err != nil || storedRefreshToken == nil {
		// save refresh token to db
		handler.refreshTokenService.CreateRefreshToken(c.Request.Context(), refreshTokenRequest)
	} else {
		// if rf token exist in db, then update it base on username
		fmt.Println("exp time ", refreshTokenExpTime)
		handler.refreshTokenService.UpdateRefreshToken(c.Request.Context(), refreshTokenRequest)
	}

	// set access token
	c.SetCookie(
		"access_token",
		accessToken,
		constants.COOKIE_DURATION,
		"/",
		"",
		false,
		true)

	// set refresh token
	c.SetCookie(
		"refresh_token",
		refreshToken,
		constants.COOKIE_DURATION,
		"/",
		"",
		false,
		true)
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.User](&entity.User{Username: user.Username}))
}

func (handler *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Request.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}

	err = authentication.VerifyToken(refreshToken.Value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}

	// extract user info
	claims := jwt.MapClaims{}
	token, err := jwt.Parse(refreshToken.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err == nil {
		claims = token.Claims.(jwt.MapClaims)
	} else {
		c.JSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}
	user := &entity.User{
		Username: claims["username"].(string),
	}

	// check refresh token in database to see if it exipre or exist
	storedRefreshToken, err := handler.refreshTokenService.FindRefreshTokenByUsername(c.Request.Context(), user.Username) // find by username
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	if storedRefreshToken.Token != token.Raw {
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "Unauthorized", Field: "refresh_token", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}

	// generate new access token
	accessTokenExpTime := time.Now().Add(constants.JWT_DURATION)
	accessToken, err := authentication.GenerateToken(user, accessTokenExpTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}

	// set access token
	c.SetCookie(
		"access_token",
		accessToken,
		constants.COOKIE_DURATION,
		"/",
		"",
		false,
		true)
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.User](&entity.User{Username: user.Username}))
}
