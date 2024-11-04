package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type RefreshTokenRepository interface {
	CreateRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error
	FindRefreshToken(ctx context.Context, tokenValue string) (*entity.RefreshToken, error)
}