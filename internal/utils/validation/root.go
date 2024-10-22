package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func GetValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("rating", ValidateRating)
		if err != nil {
			panic(err)
		}

		err = v.RegisterValidation("special_features", ValidateSpecialFeatures)
		if err != nil {
			panic(err)
		}
	}
}
