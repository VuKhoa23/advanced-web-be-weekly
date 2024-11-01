package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(c context.Context, user *entity.User) error
	GetUserByUsername(c context.Context, username string) (*entity.User, error)
}
