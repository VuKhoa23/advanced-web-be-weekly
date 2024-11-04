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

func (service *RefreshTokenService) UpdateRefreshToken(ctx context.Context, refreshTokenRequest model.RefreshTokenRequest) error {
	refreshToken := &entity.RefreshToken{
		Token: refreshTokenRequest.Token,
		Username: refreshTokenRequest.Username,
		ExpTime: refreshTokenRequest.ExpTime,
	}
	err := service.refreshTokenRepository.UpdateRefreshToken(ctx, refreshToken)
	if err != nil {
		return  err
	}
	return nil
}

func (service *RefreshTokenService) FindRefreshTokenByUsername(ctx context.Context, username string) (*entity.RefreshToken, error) {
	refreshToken, err := service.refreshTokenRepository.FindRefreshTokenByUsername(ctx, username)
	if err != nil {
		return  nil, err
	}
	return refreshToken, nil
}

func (service *RefreshTokenService) FindRefreshTokenByValue(ctx context.Context, tokenValue string) (*entity.RefreshToken, error) {
	refreshToken, err := service.refreshTokenRepository.FindRefreshTokenByValue(ctx, tokenValue)
	if err != nil {
		return  nil, err
	}
	return refreshToken, nil
}
