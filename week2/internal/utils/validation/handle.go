package validation

import (
	"encoding/json"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func BindJsonAndValidate(c *gin.Context, dest interface{}) error {
	err := c.ShouldBindJSON(&dest)

	if err != nil {
		checkErr(c, err)
	}
	return err
}

func checkErr(c *gin.Context, err error) {
	switch t := err.(type) {
	case *json.UnmarshalTypeError:
		httpErr := httpcommon.Error{
			Message: httpcommon.ErrorMessage.InvalidDataType, Code: httpcommon.ErrorResponseCode.InvalidDataType, Field: t.Field,
		}
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpErr))
		return
	case *json.SyntaxError:
		httpErr := httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest}
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpErr))
		return
	case validator.ValidationErrors:
		httpErrs := handleValidationErrors(err)
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpErrs...))
		return
	default:
		httpErr := httpcommon.Error{Message: err.Error(), Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: ""}
		c.JSON(http.StatusBadRequest, httpErr)
		return
	}
}

func handleValidationErrors(errs error) (httpErrs []httpcommon.Error) {
	for _, fieldErr := range errs.(validator.ValidationErrors) {
		errCodeWithStructNs := httpcommon.CustomValidationErrCode[strings.ToLower(fieldErr.StructNamespace())]
		if errCodeWithStructNs == "" {
			// handle builtin validation
			httpErrs = append(httpErrs, httpcommon.Error{
				Message: httpcommon.ErrorMessage.InvalidRequest, Code: httpcommon.ErrorResponseCode.InvalidRequest, Field: fieldErr.Field(),
			})
		} else {
			// handle custom validation
			httpErrs = append(httpErrs, httpcommon.Error{
				Message: httpcommon.ErrorMessage.InvalidRequest, Code: errCodeWithStructNs, Field: fieldErr.Field(),
			})
		}
	}
	return httpErrs
}
