package repositoryimplement

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db database.Db) repository.UserRepository {
	return &UserRepository{db: db}
}

// CreateUser implements repository.UserRepository.
func (repo *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	err := repo.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	result := repo.db.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
