package authentication

import (
	"fmt"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Exp      int64  `json:"exp"`
}

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateAccessToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
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
