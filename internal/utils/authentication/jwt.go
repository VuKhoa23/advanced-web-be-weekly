package authentication

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

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

func VerifyToken(tokenString string, requestUrl string, requestTime int64) error {
	now := time.Now().Unix()

	if now-requestTime > 60 {
		return fmt.Errorf("invalid request time")
	}

	token, err := GenerateTokenFromApiKey(requestUrl, requestTime)

	if err != nil {
		return err
	}

	if token != tokenString {
		return fmt.Errorf("invalid token")
	}

	return nil
}
