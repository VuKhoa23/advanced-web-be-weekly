package repository

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type UserRepository interface {
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
}
