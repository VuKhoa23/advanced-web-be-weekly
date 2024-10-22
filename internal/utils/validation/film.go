package validation

import (
	"strings"

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

func ValidateSpecialFeatures(fl validator.FieldLevel) bool {
    validFeatures := map[string]struct{}{
        "Trailers":       {},
        "Commentaries":   {},
        "Deleted Scenes": {},
        "Behind the Scenes": {},
    }

    features := fl.Field().String()
    featureList := strings.Split(features, ",")

    for _, feature := range featureList {
        if _, exists := validFeatures[feature]; !exists {
            return false
        }
    }

    return true
}

