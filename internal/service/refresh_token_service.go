package service

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type RefreshTokenService interface {
	CreateRefreshToken(ctx context.Context, refreshTokenRequest model.RefreshTokenRequest) error
	FindRefreshToken(ctx context.Context, tokenValue string) (*entity.RefreshToken, error)
}