package repositoryimplement

import (
	"context"
	"time"

	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db database.Db) repository.RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (repo *RefreshTokenRepository) CreateRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error {
	err := repo.db.WithContext(ctx).Create(refreshToken).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *RefreshTokenRepository) FindRefreshToken(ctx context.Context, tokenValue string) (*entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken
	err := repo.db.WithContext(ctx).Where("token = ?", tokenValue).First(&refreshToken).Error
	if err != nil {
		return nil, err
	}
	if refreshToken.ExpTime.Before(time.Now()) {
		return nil, nil
	}
	return &refreshToken, nil
}

