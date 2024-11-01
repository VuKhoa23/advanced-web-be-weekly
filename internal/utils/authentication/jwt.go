package authentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"os"
	"time"
)

type Claims struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Exp      int64  `json:"exp"`
}

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateAccessToken(user *entity.User) (string, error) {
	claims := Claims{
		UserName: user.UserName,
		Password: user.Password,
		Exp:      time.Now().Add(time.Hour * 24).Unix(),
	}

	// Mã hóa payload thành JSON
	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Tính toán hash
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write(payloadBytes)
	hash := h.Sum(nil)

	// Kết hợp payload và hash
	token := string(payloadBytes) + string(hash)

	// Mã hóa bằng base64
	return base64.StdEncoding.EncodeToString([]byte(token)), nil
}

func VerifyToken(tokenString string) (bool, error) {
	// Giải mã token
	tokenBytes, err := base64.StdEncoding.DecodeString(tokenString)
	if err != nil {
		return false, err
	}

	// Tách payload và hash
	payloadLength := len(tokenString) / 2
	payload := tokenBytes[:payloadLength]
	hash := tokenBytes[payloadLength:]

	// Tính toán lại hash
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write(payload)
	newHash := h.Sum(nil)

	// So sánh hash
	if !hmac.Equal(hash, newHash) {
		return false, fmt.Errorf("invalid token")
	}

	// Giải mã payload
	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return false, err
	}

	// Kiểm tra thời gian hết hạn
	if claims.Exp < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	return true, nil
}
