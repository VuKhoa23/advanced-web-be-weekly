package service

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type UserService interface {
	CheckPassword(ctx context.Context, userRequest model.UserRequest) (*entity.User, error)
	CreateUser(ctx context.Context, userRequest model.UserRequest) (string, error)
}
