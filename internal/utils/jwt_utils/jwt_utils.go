package utils

import (
	"fmt"
	"time"

	"github.com/VuKhoa23/advanced-web-be/internal/constants"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"id":       id,
			"exp":      time.Now().Add(time.Hour * 1000).Unix(),
		})

	tokenString, err := token.SignedString([]byte(constants.JWT_SERCRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_SERCRET), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	idClaim, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("id claim is missing or not a number")
	}
	id := int64(idClaim)

	return id, nil
}
