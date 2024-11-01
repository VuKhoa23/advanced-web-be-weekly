package middleware

import (
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyTokenMiddleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}

	_, err = authentication.VerifyToken(cookie.Value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}
}
