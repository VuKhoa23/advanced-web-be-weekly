package middleware

import (
	"fmt"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func VerifyTokenMiddleware(c *gin.Context) {
	requestTime, err := strconv.ParseInt(c.Request.Header.Get("Request-Time"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "Unauthorized", Field: "Request-Time", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}
	token := c.Request.Header.Get("Token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "Unauthorized", Field: "Token", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}

	fmt.Println("http://" + c.Request.Host + c.Request.URL.Path)

	// IDK if there's a more elegant way to retrieve the entire request URL in Gin
	apiPath := strings.TrimSuffix(c.Request.URL.Path, "/")
	err = authentication.VerifyToken(token, "http://"+c.Request.Host+apiPath, requestTime)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.Unauthorized,
		}))
		return
	}
}
