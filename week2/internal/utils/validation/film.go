package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidateRating(fl validator.FieldLevel) bool {
	var validRatings = []string{"G", "PG", "PG-13", "R", "NC-17"}

	rating := fl.Field().String()

	for _, validRating := range validRatings {
		if rating == validRating {
			return true
		}
	}
	return false
}
