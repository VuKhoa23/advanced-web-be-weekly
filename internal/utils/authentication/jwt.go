package authentication

import (
	"fmt"
	"os"
	"time"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(user *entity.User, expTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Username,
			"exp":      expTime.Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

var apiKey = os.Getenv("API_KEY")

func GenerateTokenFromApiKey(requestUrl string, requestTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"url":  requestUrl,
			"time": requestTime,
		})

	tokenString, err := token.SignedString([]byte(apiKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}