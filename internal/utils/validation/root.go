package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func GetValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("rating", ValidateRating)
	}
}
