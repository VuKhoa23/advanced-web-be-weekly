package serviceimplement

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type RefreshTokenService struct {
	refreshTokenRepository repository.RefreshTokenRepository
}

func NewRefreshTokenService(refreshTokenRepository repository.RefreshTokenRepository) service.RefreshTokenService {
	return &RefreshTokenService{refreshTokenRepository: refreshTokenRepository}
}

func (service *RefreshTokenService) CreateRefreshToken(ctx context.Context, refreshTokenRequest model.RefreshTokenRequest) error {
	refreshToken := &entity.RefreshToken{
		Token: refreshTokenRequest.Token,
		Username: refreshTokenRequest.Username,
		ExpTime: refreshTokenRequest.ExpTime,
	}
	err := service.refreshTokenRepository.CreateRefreshToken(ctx, refreshToken)
	if err != nil {
		return  err
	}
	return nil
}

func (service *RefreshTokenService) FindRefreshToken(ctx context.Context, tokenValue string) (*entity.RefreshToken, error) {
	refreshToken, err := service.refreshTokenRepository.FindRefreshToken(ctx, tokenValue)
	if err != nil {
		return  nil, err
	}
	return refreshToken, nil
}
